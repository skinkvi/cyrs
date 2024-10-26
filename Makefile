TERN := tern 
GO := go 

# создание новой миграции 
new-migration:
	@echo "Введите новую название новой миграции"
	@read migration_name
	@$(TERN) new $(migration_name)

# применение миграций 
migrate:
	@$(TERN) migrate

# откат миграций 
rollback:
	@$(TERN) rollback

# запуск программы
