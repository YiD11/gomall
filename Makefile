export ROOT_MOD = github.com/YiD11/gomall
.PHONY: hello
hello:
	echo hello

.PHONY: gen-frontend
gen-frontend:
	@cd ./app/frontend && cwgo server --type HTTP --idl ../.././idl/frontend/order_page.proto --service frontend -module ${ROOT_MOD}/app/frontend -I ../.././idl/

.PHONY: gen-client
gen-client:
	cwgo client --type RPC --idl ../idl/user.proto --service user -module ${ROOT_MOD}/rpc_gen -I ../idl/

.PHONY: gen-user
gen-user:
	@cd app/user && cwgo server --type RPC  --service user --module  ${ROOT_MOD}/app/user  --pass "-use  ${ROOT_MOD}/rpc_gen/kitex_gen" -I ../../idl  --idl ../../idl/user.proto
	@cd rpc_gen && cwgo client --type RPC  --service user --module  ${ROOT_MOD}/rpc_gen --I ../idl --idl ../idl/user.proto

.PHONY: gen-product
gen-product:
	@cd app/product && cwgo server --type RPC  --service product --module  ${ROOT_MOD}/app/product  --pass "-use  ${ROOT_MOD}/rpc_gen/kitex_gen" -I ../../idl  --idl ../../idl/product.proto
	@cd rpc_gen && cwgo client --type RPC  --service product --module  ${ROOT_MOD}/rpc_gen --I ../idl --idl ../idl/product.proto

.PHONY: gen-cart
gen-cart:
	@cd app/cart && cwgo server --type RPC  --service cart --module  ${ROOT_MOD}/app/cart  --pass "-use  ${ROOT_MOD}/rpc_gen/kitex_gen" -I ../../idl  --idl ../../idl/cart.proto
	@cd rpc_gen && cwgo client --type RPC  --service cart --module  ${ROOT_MOD}/rpc_gen --I ../idl --idl ../idl/cart.proto

.PHONY: gen-payment
gen-payment:
	@cd rpc_gen && cwgo client --type RPC  --service payment --module  ${ROOT_MOD}/rpc_gen --I ../idl --idl ../idl/payment.proto
	@cd app/payment && cwgo server --type RPC  --service payment --module  ${ROOT_MOD}/app/payment  --pass "-use  ${ROOT_MOD}/rpc_gen/kitex_gen" -I ../../idl  --idl ../../idl/payment.proto

.PHONY: gen-checkout
gen-checkout:
	@cd rpc_gen && cwgo client --type RPC  --service checkout --module  ${ROOT_MOD}/rpc_gen --I ../idl --idl ../idl/checkout.proto
	@cd app/checkout && cwgo server --type RPC  --service checkout --module  ${ROOT_MOD}/app/checkout  --pass "-use  ${ROOT_MOD}/rpc_gen/kitex_gen" -I ../../idl  --idl ../../idl/checkout.proto

.PHONY: gen-order
gen-order:
	@cd rpc_gen && cwgo client --type RPC  --service order --module  ${ROOT_MOD}/rpc_gen --I ../idl --idl ../idl/order.proto
	@cd app/order && cwgo server --type RPC  --service order --module  ${ROOT_MOD}/app/order  --pass "-use  ${ROOT_MOD}/rpc_gen/kitex_gen" -I ../../idl  --idl ../../idl/order.proto

.PHONY: gen-email
gen-email:
	@cd rpc_gen && cwgo client --type RPC  --service email --module  ${ROOT_MOD}/rpc_gen --I ../idl --idl ../idl/email.proto
	@cd app/email && cwgo server --type RPC  --service email --module  ${ROOT_MOD}/app/email  --pass "-use  ${ROOT_MOD}/rpc_gen/kitex_gen" -I ../../idl  --idl ../../idl/email.proto

.PHONY: gen-test
gen-test:
	@cd rpc_gen && cwgo client --type RPC  --service test --module  ${ROOT_MOD}/rpc_gen --I ../idl --idl ../idl/test.proto
	@cd app/test && cwgo server --type RPC  --service test --module  ${ROOT_MOD}/app/test  --pass "-use  ${ROOT_MOD}/rpc_gen/kitex_gen" -I ../../idl  --idl ../../idl/test.proto

.PHONY: build-frontend
build-frontend:
	docker build -f ./deploy/frontend.Dockerfile -t gomall_frontend:${v} .

.PHONY: run-frontend
run-frontend:
	docker run -it --rm --name frontend --network gomall_default --env-file=./app/frontend/.env -p 8090:8090 gomall_frontend:${v}

.PHONY: build-svc
build-svc:
	docker build -f ./deploy/svc.Dockerfile -t gomall_${svc}:${v} --build-arg SVC=${svc} .

.PHONY: run-svc
run-svc:
	docker run -it --rm --name ${svc} -v ./app/${svc}/conf:/opt/gomall/${svc}/conf --network gomall_default --env-file=./app/${svc}/.env gomall_${svc}:${v}

.PHONY: tidy-all
tidy-all:
	@cd app/user && go mod tidy
	@cd app/product && go mod tidy
	@cd app/cart && go mod tidy
	@cd app/payment && go mod tidy
	@cd app/checkout && go mod tidy
	@cd app/order && go mod tidy
	@cd app/email && go mod tidy
	@cd app/frontend && go mod tidy

.PHONY: build-all
build-all:
	docker build -f ./deploy/frontend.Dockerfile -t frontend:${v} . 
	docker build -f ./deploy/svc.Dockerfile -t user:${v} --build-arg SVC=user . 
	docker build -f ./deploy/svc.Dockerfile -t product:${v} --build-arg SVC=product . 
	docker build -f ./deploy/svc.Dockerfile -t checkout:${v} --build-arg SVC=checkout . 
	docker build -f ./deploy/svc.Dockerfile -t cart:${v} --build-arg SVC=cart . 
	docker build -f ./deploy/svc.Dockerfile -t email:${v} --build-arg SVC=email . 
	docker build -f ./deploy/svc.Dockerfile -t order:${v} --build-arg SVC=order . 
	docker build -f ./deploy/svc.Dockerfile -t payment:${v} --build-arg SVC=payment . 

.PHONY: cs ce
cs:
	@cd cluster && docker compose up -d
ce:
	@cd cluster && docker compose down