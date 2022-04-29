package main

import (
	sap_api_caller "sap-api-integrations-planned-order-reads/SAP_API_Caller"
	"sap-api-integrations-planned-order-reads/SAP_API_Input_Reader"

	"github.com/latonaio/golang-logging-library-for-sap/logger"
)

func main() {
	l := logger.NewLogger()
	fr := sap_api_input_reader.NewFileReader()
	inoutSDC := fr.ReadSDC("./Inputs//SDC_Planned_Order_Component_Material_Plant_sample.json")
	caller := sap_api_caller.NewSAPAPICaller(
		"https://sandbox.api.sap.com/s4hanacloud/sap/opu/odata4/sap/", l,
	)

	accepter := inoutSDC.Accepter
	if len(accepter) == 0 || accepter[0] == "All" {
		accepter = []string{
			"Header", "HeaderMaterialPlant", "ComponentMaterialPlant",
		}
	}

	caller.AsyncGetPlannedOrder(
		inoutSDC.PlannedOrder.PlannedOrder,
		inoutSDC.PlannedOrder.Material,
		inoutSDC.PlannedOrder.MRPPlant,
		inoutSDC.PlannedOrder.Component.Plant,
		accepter,
	)
}
