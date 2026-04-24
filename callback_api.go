package antilopay

import (
	"crypto"
	"crypto/rsa"
	"crypto/sha256"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	"github.com/gofiber/fiber/v3"
)

func (c *Client) VerifySignature(body []byte, signature string) error {
	sigBytes, err := base64.StdEncoding.DecodeString(signature)
	if err != nil {
		return fmt.Errorf("decode signature: %w", err)
	}

	hash := sha256.New()
	hash.Write(body)
	hashed := hash.Sum(nil)

	return rsa.VerifyPKCS1v15(c.callbackPublicKey, crypto.SHA256, hashed, sigBytes)
}

func (c *Client) WebhookHandlerFiber(processor func(p *CallbackPayload) error) fiber.Handler {
	return func(cF fiber.Ctx) error {
		signature := cF.Get("X-Apay-Callback")
		if signature == "" {
			return cF.SendStatus(fiber.StatusForbidden)
		}

		body := cF.Body()
		if err := c.VerifySignature(body, signature); err != nil {
			return cF.SendStatus(fiber.StatusUnauthorized)
		}

		var payload CallbackPayload
		if err := json.Unmarshal(body, &payload); err != nil {
			return cF.SendStatus(fiber.StatusBadRequest)
		}

		if err := processor(&payload); err != nil {
			// Return ONLY if errored during processing
			return cF.SendStatus(fiber.StatusInternalServerError)
		}

		return cF.SendStatus(fiber.StatusOK)
	}
}

func (c *Client) WebhookHandlerHTTP(publicKey string, processor func(p *CallbackPayload) error) http.HandlerFunc {
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

		if err = c.VerifySignature(body, signature); err != nil {
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
