# User Mocks
mockgen -destination=mocks/mock_user_application.go -package=mocks --build_flags=--mod=mod github.com/solrac97gr/go-test-tables/users/domain/ports Application &&
mockgen -destination=mocks/mock_user_repository.go -package=mocks --build_flags=--mod=mod github.com/solrac97gr/go-test-tables/users/domain/ports Repository &&
mockgen -destination=mocks/mock_validator.go -package=mocks --build_flags=--mod=mod github.com/solrac97gr/go-test-tables/validator/domain/ports Validator &&
echo "Mocks generated successfully ✨"