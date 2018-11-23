project_id := ${PROJECT_ID}
version := ${GAE_VERSION}

serve: ## Execute local server
	go run main.go

deploy:
	goapp deploy -application ${project_id} -version ${version} app
