package main

func mergeMaps(map1, map2 map[string]interface{}) map[string]interface{} {
	mergedMap := make(map[string]interface{})

	// Copiar elementos do primeiro mapa para o mapa combinado
	for key, value := range map1 {
		mergedMap[key] = value
	}

	// Copiar elementos do segundo mapa para o mapa combinado (sobrescrevendo valores duplicados)
	for key, value := range map2 {
		mergedMap[key] = value
	}

	return mergedMap
}
