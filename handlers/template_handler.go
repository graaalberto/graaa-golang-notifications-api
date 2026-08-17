package handlers

import (
	"fmt"
)

// BuildMessageFromTemplate formata a mensagem com base no tipo
func BuildMessageFromTemplate(msgType string, meta map[string]interface{}) string {
	switch msgType {
	case "trip_assigned":
		return fmt.Sprintf("🚖 *FrotaGo Angola - Nova Corrida Atribuída*\nOlá! A sua viatura (%v) tem uma nova corrida atribuída.\n*Passageiro:* %v (%v)\n*Origem:* %v\n*Destino:* %v\n*Valor:* %v AOA",
			meta["plate"], meta["passengerName"], meta["passengerPhone"], meta["origin"], meta["destination"], meta["fareAOA"])

	case "trip_completed":
		return fmt.Sprintf("✅ *FrotaGo Angola - Corrida Concluída*\nObrigado por viajar connosco!\n*Valor Pago:* %v AOA\n*Recibo:* %v",
			meta["amountAOA"], meta["receiptUrl"])

	case "alert_critical":
		return fmt.Sprintf("🚨 *ALERTA CRÍTICO DE FROTA*\nViatura: %v\nAlerta: %v\nCoordenadas: %v, %v\nPor favor, contacte o despachante imediatamente!",
			meta["plate"], meta["alertMessage"], meta["lat"], meta["lng"])

	case "breakdown_alert":
		return fmt.Sprintf("🛠️ *FROTAGO MANUTENÇÃO - NOVA AVARIA*\nViatura: %v\nGravidade: %v\nProblema: %v\nOficina designada: %v",
			meta["plate"], meta["severity"], meta["description"], meta["workshop"])

	default:
		if custom, ok := meta["text"].(string); ok {
			return custom
		}
		return "FrotaGo Angola: Tem uma nova atualização no seu aplicativo."
	}
}
