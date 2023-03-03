mongoimport \
	--db 'movil_parcial' \
	--collection 'users' \
	--file '/data/users.json' \
	--jsonArray \
	--uri "mongodb://root:dev@mongo:27017" \
	--authenticationDatabase 'admin'

mongoimport \
	--db 'movil_parcial' \
	--collection 'products' \
	--file '/data/products.json' \
	--jsonArray \
	--uri "mongodb://root:dev@mongo:27017" \
	--authenticationDatabase 'admin'