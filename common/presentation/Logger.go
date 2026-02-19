package presentation

import (
	"fmt"
	"os"
	"time"

	"github.com/gofiber/fiber/v2"
)

func LoggerMiddleware(c *fiber.Ctx) error {
	start := time.Now()

	err := c.Next()

	duration := time.Since(start)

	raw := BuildRawHttpRequest(c, duration)

	fmt.Println(raw)

	file, errFile := os.OpenFile(
		"request.log",
		os.O_CREATE|os.O_WRONLY|os.O_APPEND,
		0644,
	)
	if errFile != nil {
		fmt.Println("failed open log file:", errFile)
		return err
	}
	defer file.Close()

	// Write log
	if _, errWrite := file.WriteString(raw); errWrite != nil {
		fmt.Println("failed write log:", errWrite)
	}

	return err
}

func BuildRawHttpRequest(c *fiber.Ctx, duration time.Duration) string {
	req := c.Request()

	// Date-time ISO-8601
	now := time.Now().Format("2006-01-02T15:04:05-07:00")

	// 1. Request Line
	requestLine := fmt.Sprintf("%s %s HTTP/1.1",
		req.Header.Method(),
		req.URI().RequestURI(),
	)

	// 2. Raw Headers
	rawHeaders := ""
	req.Header.VisitAll(func(k, v []byte) {
		rawHeaders += fmt.Sprintf("%s: %s\n", k, v)
	})

	// 3. Raw Body (data apa pun termasuk multipart & binary)
	rawBody := string(req.Body())

	// 4. Gabungkan
	return fmt.Sprintf(
		`================ %s ================
IP: %s
Took: %s

%s
%s
%s
=============================================
`,
		now, // <-- DATE TIME
		c.IP(),
		duration,
		requestLine, // GET /api x
		rawHeaders,  // headers
		rawBody,     // body (raw multipart)
	)
}
