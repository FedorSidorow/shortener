// Этот анализатор состоит из:
// стандартных статических анализаторов пакета golang.org/x/tools/go/analysis/passes
// всех анализаторов класса SA пакета staticcheck.io;
// двух публичных анализаторов "github.com/gordonklaus/ineffassign/pkg/ineffassign"
// и "github.com/timakin/bodyclose/passes/bodyclose"
//
// Для запуска, в корневой папке проекта, используйте команду: go run cmd/staticlint/staticlint.go ./...

package main

import (
	"strings"

	exitcheker "github.com/FedorSidorow/shortener/internal/staticlint"
	"github.com/gordonklaus/ineffassign/pkg/ineffassign"
	"github.com/timakin/bodyclose/passes/bodyclose"
	"golang.org/x/tools/go/analysis"
	"golang.org/x/tools/go/analysis/multichecker"
	"golang.org/x/tools/go/analysis/passes/printf"
	"golang.org/x/tools/go/analysis/passes/shadow"
	"golang.org/x/tools/go/analysis/passes/structtag"
	"honnef.co/go/tools/staticcheck"
	"honnef.co/go/tools/stylecheck"
)

// Добавление всех анализаторов класса SA пакета staticcheck.io.
func getStaticcheckAnalyzersSA(checks []*analysis.Analyzer) []*analysis.Analyzer {

	for _, v := range staticcheck.Analyzers {
		if strings.HasPrefix(v.Analyzer.Name, "SA") {
			checks = append(checks, v.Analyzer)
		}
	}
	return checks
}

// Добавление анализатора остальных классов пакета staticcheck.io.
func getStylecheckAnalyzers(checks []*analysis.Analyzer, data map[string]bool) []*analysis.Analyzer {

	for _, v := range stylecheck.Analyzers {

		if data[v.Analyzer.Name] {
			checks = append(checks, v.Analyzer)
		}
	}
	return checks
}

// Проверка с помощью выбранных анализаторов.
func main() {
	var checks []*analysis.Analyzer

	checks = getStaticcheckAnalyzersSA(checks)
	checks = getStylecheckAnalyzers(checks, map[string]bool{"ST1001": true})

	checks = append(checks, printf.Analyzer)
	checks = append(checks, shadow.Analyzer)
	checks = append(checks, structtag.Analyzer)

	// Два публичных анализатора кода.
	checks = append(checks, ineffassign.Analyzer)
	checks = append(checks, bodyclose.Analyzer)

	// Собственный анализатор, запрещающий использовать прямой вызов os.Exit в функции main пакета main.
	checks = append(checks, exitcheker.ExitAnalyzer)

	multichecker.Main(
		checks...,
	)
}
