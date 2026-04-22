package main

import (
	"crypto"
	"crypto/rsa"
	"crypto/sha256"
	"crypto/x509"
	"encoding/base64"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"

	"github.com/gofiber/fiber/v3"
)

func VerifySignature(body []byte, signature, publicKeyBase64 string) error {
	pubBytes, err := base64.StdEncoding.DecodeString(publicKeyBase64)
	if err != nil {
		return fmt.Errorf("decode public key: %w", err)
	}

	pubKeyInterface, err := x509.ParsePKIXPublicKey(pubBytes)
	if err != nil {
		return fmt.Errorf("parse public key: %w", err)
	}

	pubKey, ok := pubKeyInterface.(*rsa.PublicKey)
	if !ok {
		return errors.New("not an RSA public key")
	}

	sigBytes, err := base64.StdEncoding.DecodeString(signature)
	if err != nil {
		return fmt.Errorf("decode signature: %w", err)
	}

	hash := sha256.New()
	hash.Write(body)
	hashed := hash.Sum(nil)

	return rsa.VerifyPKCS1v15(pubKey, crypto.SHA256, hashed, sigBytes)
}

func WebhookHandlerFiber(publicKey string, processor func(p *CallbackPayload) error) fiber.Handler {
	return func(c fiber.Ctx) error {
		signature := c.Get("X-Apay-Callback")
		if signature == "" {
			return c.SendStatus(fiber.StatusForbidden)
		}

		body := c.Body()
		if err := VerifySignature(body, signature, publicKey); err != nil {
			return c.SendStatus(fiber.StatusUnauthorized)
		}

		var payload CallbackPayload
		if err := json.Unmarshal(body, &payload); err != nil {
			return c.SendStatus(fiber.StatusBadRequest)
		}

		if err := processor(&payload); err != nil {
			// Return ONLY if errored during processing
			return c.SendStatus(fiber.StatusInternalServerError)
		}

		return c.SendStatus(fiber.StatusOK)
	}
}

func WebhookHandlerHTTP(publicKey string, processor func(p *CallbackPayload) error) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodPost {
			w.WriteHeader(http.StatusMethodNotAllowed)
			return
		}

		signature := r.Header.Get("X-Apay-Callback")
		if signature == "" {
			w.WriteHeader(http.StatusForbidden)
			return
		}

		body, err := io.ReadAll(r.Body)
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			return
		}
		defer r.Body.Close()

		if err = VerifySignature(body, signature, publicKey); err != nil {
			w.WriteHeader(http.StatusUnauthorized)
			return
		}

		var payload CallbackPayload
		if err = json.Unmarshal(body, &payload); err != nil {
			w.WriteHeader(http.StatusBadRequest)
			return
		}

		if err = processor(&payload); err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
		w.WriteHeader(http.StatusOK)
	}
}

// Example of "Processing" the response on client
//processOrder := func(p *antilopay.CallbackPayload) error {
//	switch p.Type {
//	case "payment":
//		if p.Status == "SUCCESS" {
//			fmt.Printf("Order %s paid!\n", p.OrderID)
//		}
//	case "topup":
//		fmt.Printf("Steam Topup %s status: %s\n", p.TopupID, p.Status)
//	}
//	return nil
//}
