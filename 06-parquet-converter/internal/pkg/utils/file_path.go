package utils

// func GetCsvFilePath() string {
// 	// Obtém o diretório de trabalho atual
// 	cwd, err := os.Getwd()
// 	if err != nil {
// 		fmt.Println("Erro ao obter o diretório de trabalho atual:", err)
// 	}
// 	// Volta uma pasta e entra em /internal/pkg/data/test/
// 	newPath := filepath.Join(filepath.Dir(cwd), "/pkg/data/csv/")

// 	fmt.Println(newPath)
// 	return newPath
// }
// func GetJsonFilePath() string {
// 	// Obtém o diretório de trabalho atual
// 	cwd, err := os.Getwd()
// 	if err != nil {
// 		fmt.Println("Erro ao obter o diretório de trabalho atual:", err)
// 	}
// 	// Volta uma pasta e entra em /internal/pkg/data/test/
// 	newPath := filepath.Join(filepath.Dir(cwd), "/pkg/data/json/")

// 	fmt.Println(newPath)
// 	return newPath
// }
// func GetParquetFilePath() string {
// 	// Obtém o diretório de trabalho atual
// 	cwd, err := os.Getwd()
// 	if err != nil {
// 		fmt.Println("Erro ao obter o diretório de trabalho atual:", err)
// 	}
// 	// Volta uma pasta e entra em /internal/pkg/data/test/
// 	newPath := filepath.Join(filepath.Dir(cwd), "/pkg/data/parquet/")

// 	fmt.Println(newPath)
// 	return newPath
// }
// func GetTestFilePath() string {
// 	// Obtém o diretório de trabalho atual
// 	cwd, err := os.Getwd()
// 	if err != nil {
// 		fmt.Println("Erro ao obter o diretório de trabalho atual:", err)
// 	}
// 	// Volta uma pasta e entra em /internal/pkg/data/test/
// 	newPath := filepath.Join(filepath.Dir(cwd), "/pkg/data/test/")

// 	fmt.Println(newPath)
// 	return newPath
// }

func GetCsvFilePath() string {
	// absPath, err := filepath.Abs("../../.env")
	// if err != nil {
	// 	log.Fatalf("Error finding absolute path: %v", err)
	// }
	// fmt.Println("Absolute path:", absPath)

	// err = godotenv.Load("../../.env")
	// if err != nil {
	// 	log.Fatalf("Error loading .env file")
	// }
	// return os.Getenv("COMMON_PATH") + "golang-studies/06-parquet-converter/internal/pkg/data/csv/"
	return "/home/pedro/workspace/studies/golang-studies/06-parquet-converter/internal/pkg/data/csv/"
}
func GetJsonFilePath() string {
	// err := godotenv.Load("../../.env")
	// if err != nil {
	// 	log.Fatalf("Error loading .env file")
	// }
	// return os.Getenv("COMMON_PATH") + "golang-studies/06-parquet-converter/internal/pkg/data/json/"
	return "/home/pedro/workspace/studies/golang-studies/06-parquet-converter/internal/pkg/data/json/"
}
func GetParquetFilePath() string {
	// err := godotenv.Load("../../.env")
	// if err != nil {
	// 	log.Fatalf("Error loading .env file")
	// }
	// return os.Getenv("COMMON_PATH") + "golang-studies/06-parquet-converter/internal/pkg/data/parquet/"
	return "/home/pedro/workspace/studies/golang-studies/06-parquet-converter/internal/pkg/data/parquet/"
}
func GetTestFilePath() string {
	// err := godotenv.Load("../../.env")
	// if err != nil {
	// 	log.Fatalf("Error loading .env file")
	// }
	// return os.Getenv("COMMON_PATH") + "golang-studies/06-parquet-converter/internal/pkg/data/test/"
	return "/home/pedro/workspace/studies/golang-studies/06-parquet-converter/internal/pkg/data/test/"
}
