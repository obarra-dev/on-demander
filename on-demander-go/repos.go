package main

func getRepos() []string {
	return []string{
		`{"host":"github.com","repoOwner":"obarra-dev","repoName":"dependency-track","branch":"master","generatePrToFixIssues":false}`,
		`{"host":"github.com","repoOwner":"obarra-dev","repoName":"dependency-track","branch":"source2","generatePrToFixIssues":false}`,
		`{"host":"github.com","repoOwner":"obarra-dev","repoName":"dependency-track","branch":"master","generatePrToFixIssues":false}`,
		`{"host":"github.com","repoOwner":"obarra-dev","repoName":"dependency-track","branch":"source2","generatePrToFixIssues":false}`,
		`{"host":"github.com","repoOwner":"obarra-dev","repoName":"dependency-track","branch":"master","generatePrToFixIssues":false}`,
		`{"host":"github.com","repoOwner":"obarra-dev","repoName":"dependency-track","branch":"source2","generatePrToFixIssues":false}`,
		`{"host":"github.com","repoOwner":"obarra-dev","repoName":"dependency-track","branch":"master","generatePrToFixIssues":false}`,
		`{"host":"github.com","repoOwner":"obarra-dev","repoName":"dependency-track","branch":"source2","generatePrToFixIssues":false}`,
		`{"host":"github.com","repoOwner":"obarra-dev","repoName":"dependency-track","branch":"master","generatePrToFixIssues":false}`,
		`{"host":"github.com","repoOwner":"obarra-dev","repoName":"dependency-track","branch":"source2","generatePrToFixIssues":false}`,
		`{"host":"github.com","repoOwner":"obarra-dev","repoName":"dependency-track","branch":"master","generatePrToFixIssues":false}`,
		`{"host":"github.com","repoOwner":"obarra-dev","repoName":"dependency-track","branch":"source2","generatePrToFixIssues":false}`,
		`{"host":"github.com","repoOwner":"obarra-dev","repoName":"dependency-track","branch":"master","generatePrToFixIssues":false}`,
		`{"host":"github.com","repoOwner":"obarra-dev","repoName":"dependency-track","branch":"source2","generatePrToFixIssues":false}`,
		`{"host":"github.com","repoOwner":"obarra-dev","repoName":"dependency-track","branch":"master","generatePrToFixIssues":false}`,
		`{"host":"github.com","repoOwner":"obarra-dev","repoName":"dependency-track","branch":"source2","generatePrToFixIssues":false}`,

		/*
			`{"host":"github.com","repoOwner":"obarra-dev","repoName":"minesweeper-API","branch":"add-errors","generatePrToFixIssues":false}`,
			`{"host":"github.com","repoOwner":"obarra-dev","repoName":"minesweeper-API","branch":"main","generatePrToFixIssues":false}`,

			`{"host":"github.com","repoOwner":"obarra-dev","repoName":"errorprone-local-test","branch":"errorprone-2.18.0","generatePrToFixIssues":false}`,
			`{"host":"github.com","repoOwner":"obarra-dev","repoName":"errorprone-local-test","branch":"main","generatePrToFixIssues":false}`,

			`{"host":"github.com","repoOwner":"obarra-dev","repoName":"BIper","branch":"master","generatePrToFixIssues":false}`,

			`{"host":"github.com","repoOwner":"obarra-dev","repoName":"dependency-track","branch":"master","generatePrToFixIssues":false}`,
			`{"host":"github.com","repoOwner":"obarra-dev","repoName":"dependency-track","branch":"source2","generatePrToFixIssues":false}`,

			`{"host":"github.com","repoOwner":"obarra-dev","repoName":"forecast-service","branch":"master","generatePrToFixIssues":false}`,
			`{"host":"github.com","repoOwner":"obarra-dev","repoName":"forecast-service","branch":"generator-logic","generatePrToFixIssues":false}`,

			`{"host":"github.com","repoOwner":"obarra-dev","repoName":"minesweeper-java-client","branch":"main","generatePrToFixIssues":false}`,

			`{"host":"github.com","repoOwner":"obarra-dev","repoName":"TestRepo","branch":"main","generatePrToFixIssues":false}`,
			`{"host":"github.com","repoOwner":"obarra-dev","repoName":"TestRepo","branch":"test-pom","generatePrToFixIssues":false}`,

			`{"host":"github.com","repoOwner":"obarra-dev","repoName":"minesweeper-API","branch":"main","generatePrToFixIssues":false}`,
			`{"host":"github.com","repoOwner":"obarra-dev","repoName":"minesweeper-API","branch":"add-errors","generatePrToFixIssues":false}`,

			`{"host":"github.com","repoOwner":"obarra-dev","repoName":"errorprone-local-test","branch":"errorprone-2.18.0","generatePrToFixIssues":false}`,
			`{"host":"github.com","repoOwner":"obarra-dev","repoName":"errorprone-local-test","branch":"main","generatePrToFixIssues":false}`,

			`{"host":"github.com","repoOwner":"obarra-dev","repoName":"BIper","branch":"master","generatePrToFixIssues":false}`,

			`{"host":"github.com","repoOwner":"obarra-dev","repoName":"dependency-track","branch":"master","generatePrToFixIssues":false}`,
			`{"host":"github.com","repoOwner":"obarra-dev","repoName":"dependency-track","branch":"source2","generatePrToFixIssues":false}`,

			`{"host":"github.com","repoOwner":"obarra-dev","repoName":"forecast-service","branch":"master","generatePrToFixIssues":false}`,
			`{"host":"github.com","repoOwner":"obarra-dev","repoName":"forecast-service","branch":"generator-logic","generatePrToFixIssues":false}`,

			`{"host":"github.com","repoOwner":"obarra-dev","repoName":"minesweeper-java-client","branch":"main","generatePrToFixIssues":false}`,

			`{"host":"github.com","repoOwner":"obarra-dev","repoName":"TestRepo","branch":"main","generatePrToFixIssues":false}`,
			`{"host":"github.com","repoOwner":"obarra-dev","repoName":"TestRepo","branch":"test-pom","generatePrToFixIssues":false}`,

			`{"host":"github.com","repoOwner":"obarra-dev","repoName":"minesweeper-API","branch":"main","generatePrToFixIssues":false}`,
			`{"host":"github.com","repoOwner":"obarra-dev","repoName":"minesweeper-API","branch":"add-errors","generatePrToFixIssues":false}`,

			`{"host":"github.com","repoOwner":"obarra-dev","repoName":"errorprone-local-test","branch":"errorprone-2.18.0","generatePrToFixIssues":false}`,
			`{"host":"github.com","repoOwner":"obarra-dev","repoName":"errorprone-local-test","branch":"main","generatePrToFixIssues":false}`,

			`{"host":"github.com","repoOwner":"obarra-dev","repoName":"BIper","branch":"master","generatePrToFixIssues":false}`,

			`{"host":"github.com","repoOwner":"obarra-dev","repoName":"dependency-track","branch":"master","generatePrToFixIssues":false}`,
			`{"host":"github.com","repoOwner":"obarra-dev","repoName":"dependency-track","branch":"source2","generatePrToFixIssues":false}`,

			`{"host":"github.com","repoOwner":"obarra-dev","repoName":"forecast-service","branch":"master","generatePrToFixIssues":false}`,
			`{"host":"github.com","repoOwner":"obarra-dev","repoName":"forecast-service","branch":"generator-logic","generatePrToFixIssues":false}`,

			`{"host":"github.com","repoOwner":"obarra-dev","repoName":"minesweeper-java-client","branch":"main","generatePrToFixIssues":false}`,

			`{"host":"github.com","repoOwner":"obarra-dev","repoName":"TestRepo","branch":"main","generatePrToFixIssues":false}`,
			`{"host":"github.com","repoOwner":"obarra-dev","repoName":"TestRepo","branch":"test-pom","generatePrToFixIssues":false}`,

			`{"host":"github.com","repoOwner":"obarra-dev","repoName":"minesweeper-API","branch":"main","generatePrToFixIssues":false}`,
			`{"host":"github.com","repoOwner":"obarra-dev","repoName":"minesweeper-API","branch":"add-errors","generatePrToFixIssues":false}`,

			`{"host":"github.com","repoOwner":"obarra-dev","repoName":"errorprone-local-test","branch":"errorprone-2.18.0","generatePrToFixIssues":false}`,
			`{"host":"github.com","repoOwner":"obarra-dev","repoName":"errorprone-local-test","branch":"main","generatePrToFixIssues":false}`,

			`{"host":"github.com","repoOwner":"obarra-dev","repoName":"BIper","branch":"master","generatePrToFixIssues":false}`,

			`{"host":"github.com","repoOwner":"obarra-dev","repoName":"dependency-track","branch":"master","generatePrToFixIssues":false}`,
			`{"host":"github.com","repoOwner":"obarra-dev","repoName":"dependency-track","branch":"source2","generatePrToFixIssues":false}`,

			`{"host":"github.com","repoOwner":"obarra-dev","repoName":"forecast-service","branch":"master","generatePrToFixIssues":false}`,
			`{"host":"github.com","repoOwner":"obarra-dev","repoName":"forecast-service","branch":"generator-logic","generatePrToFixIssues":false}`,

			`{"host":"github.com","repoOwner":"obarra-dev","repoName":"minesweeper-java-client","branch":"main","generatePrToFixIssues":false}`,

			`{"host":"github.com","repoOwner":"obarra-dev","repoName":"TestRepo","branch":"main","generatePrToFixIssues":false}`,
			`{"host":"github.com","repoOwner":"obarra-dev","repoName":"TestRepo","branch":"test-pom","generatePrToFixIssues":false}`,

			`{"host":"github.com","repoOwner":"obarra-dev","repoName":"minesweeper-API","branch":"main","generatePrToFixIssues":false}`,
			`{"host":"github.com","repoOwner":"obarra-dev","repoName":"minesweeper-API","branch":"add-errors","generatePrToFixIssues":false}`,

			`{"host":"github.com","repoOwner":"obarra-dev","repoName":"errorprone-local-test","branch":"errorprone-2.18.0","generatePrToFixIssues":false}`,
			`{"host":"github.com","repoOwner":"obarra-dev","repoName":"errorprone-local-test","branch":"main","generatePrToFixIssues":false}`,

			`{"host":"github.com","repoOwner":"obarra-dev","repoName":"BIper","branch":"master","generatePrToFixIssues":false}`,

			`{"host":"github.com","repoOwner":"obarra-dev","repoName":"dependency-track","branch":"master","generatePrToFixIssues":false}`,
			`{"host":"github.com","repoOwner":"obarra-dev","repoName":"dependency-track","branch":"source2","generatePrToFixIssues":false}`,

			`{"host":"github.com","repoOwner":"obarra-dev","repoName":"forecast-service","branch":"master","generatePrToFixIssues":false}`,
			`{"host":"github.com","repoOwner":"obarra-dev","repoName":"forecast-service","branch":"generator-logic","generatePrToFixIssues":false}`,

			`{"host":"github.com","repoOwner":"obarra-dev","repoName":"minesweeper-java-client","branch":"main","generatePrToFixIssues":false}`,

			`{"host":"github.com","repoOwner":"obarra-dev","repoName":"TestRepo","branch":"main","generatePrToFixIssues":false}`,
			`{"host":"github.com","repoOwner":"obarra-dev","repoName":"TestRepo","branch":"test-pom","generatePrToFixIssues":false}`,

			`{"host":"github.com","repoOwner":"obarra-dev","repoName":"minesweeper-API","branch":"main","generatePrToFixIssues":false}`,
			`{"host":"github.com","repoOwner":"obarra-dev","repoName":"minesweeper-API","branch":"add-errors","generatePrToFixIssues":false}`,

			`{"host":"github.com","repoOwner":"obarra-dev","repoName":"errorprone-local-test","branch":"errorprone-2.18.0","generatePrToFixIssues":false}`,
			`{"host":"github.com","repoOwner":"obarra-dev","repoName":"errorprone-local-test","branch":"main","generatePrToFixIssues":false}`,

			`{"host":"github.com","repoOwner":"obarra-dev","repoName":"BIper","branch":"master","generatePrToFixIssues":false}`,

			`{"host":"github.com","repoOwner":"obarra-dev","repoName":"dependency-track","branch":"master","generatePrToFixIssues":false}`,
			`{"host":"github.com","repoOwner":"obarra-dev","repoName":"dependency-track","branch":"source2","generatePrToFixIssues":false}`,

			`{"host":"github.com","repoOwner":"obarra-dev","repoName":"forecast-service","branch":"master","generatePrToFixIssues":false}`,
			`{"host":"github.com","repoOwner":"obarra-dev","repoName":"forecast-service","branch":"generator-logic","generatePrToFixIssues":false}`,

			`{"host":"github.com","repoOwner":"obarra-dev","repoName":"minesweeper-java-client","branch":"main","generatePrToFixIssues":false}`,

			`{"host":"github.com","repoOwner":"obarra-dev","repoName":"TestRepo","branch":"main","generatePrToFixIssues":false}`,
			`{"host":"github.com","repoOwner":"obarra-dev","repoName":"TestRepo","branch":"test-pom","generatePrToFixIssues":false}`,
		*/
	}
}
