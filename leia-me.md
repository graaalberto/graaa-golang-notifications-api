API - Notificações, WhatsApp & SMS (graaa-golang-notifications-api)

Focada no envio em tempo real de mensagens e alertas para o contexto de Angola:

    Notificações via WhatsApp (ex: confirmação de corrida ao passageiro e alerta de novo despacho ao motorista).

    SMS Gateway (Unitel / Movicel para motoristas sem dados móveis ou em zonas de baixa cobertura).

    Push Notifications & Alertas Críticos de Manutenção (para a equipa de mecânicos quando o motor sobreaquecer ou o óleo estiver degradado).

    Histórico de Mensagens gravado no fleet_db

    
 📁 Estrutura da Pasta: backend-notifications-golang

backend-notifications-golang/
├── .env
├── go.mod
├── main.go
├── config/
│   └── database.go          # Conexão com o PostgreSQL (fleet_db)
├── models/
│   └── notification.go      # Structs de mensagens enviadas e templates
├── services/
│   ├── whatsapp_service.go  # Provedor WhatsApp Business / Evolution API
│   └── sms_service.go       # Provedor SMS Angola (Unitel / Movicel / Twilio)
└── handlers/
    ├── notification_handler.go # Endpoints de disparo de SMS/WhatsApp
    └── template_handler.go     # Templates de mensagens (corrida, avaria, alerta)

🚀 Como Executar no PowerShell:

    Crie a pasta backend-notifications-golang no seu desktop/projeto e cole os arquivos.

No PowerShell:

cd backend-notifications-golang
go mod tidy
go run main.go