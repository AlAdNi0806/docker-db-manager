.PHONY: build install

build:
	@go build -o manage-db . && ./manage-db

install:
	@go build -o manage-db .
	@sudo install -m 755 manage-db /usr/local/bin/manage-db
	@echo "✅ Installed manage-db to /usr/local/bin"
	@echo "You can now run 'manage-db' from anywhere!"

uninstall:
	@sudo rm -f /usr/local/bin/manage-db
	@echo "🗑️  Removed manage-db"
