
# 🧠 AI Chat Bot using Golang

This is a simple AI-powered chatbot built with **Go**, integrated with Slack, Wit.ai, and Wolfram Alpha to respond smartly in real-time.

---

## 🚀 Features

- 🤖 Slack bot integration
- 🧠 Natural Language Processing using Wit.ai
- 📊 Knowledge-based Q&A via Wolfram Alpha
- 💬 Real-time message handling
- 🔧 Modular Go codebase

---

## 🛠️ Tech Stack

- **Go (Golang)** – main programming language
- **Slack API** – for chat interface
- **Wit.ai** – for NLP processing
- **Wolfram Alpha** – for computation & facts

---

## 📁 Folder Structure

. ├── main.go # Entry point ├── go.mod # Go module definitions ├── go.sum # Dependency checksums ├── .env # Environment variables (ignored in git) ├── .gitignore # Git ignore rules

---

## 🔐 Environment Variables

Create a `.env` file and add the following:

```env
SLACK_BOT_TOKEN="your-slack-bot-token"
SLACK_APP_TOKEN="your-slack-app-token"
WIT_AI_TOKEN="your-wit-ai-token"
WOLFRAM_APP_ID="your-wolfram-app-id"
⚠️ Do not commit .env to your repository. It must be listed in .gitignore.

📦 Installation & Running
Clone this repo:

git clone https://github.com/22CB006/AI_chat_bot-golang.git
cd AI_chat_bot-golang

Download dependencies:
go mod tidy

Run the bot:
go run main.go

🌱 Future Enhancements
Context-aware responses

Multi-platform support (Telegram, WhatsApp, etc.)

Voice input recognition

👨‍💻 Author
Arya Lakshmi
GitHub
