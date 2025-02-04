SHORT_SHA=$(git rev-parse --short=7 HEAD)
./.deploy/build_deployment.sh \
  STACK="STACK_VALUE" \
  DEPLOYMENT_ENVIROMENT="PROD" \
  IMAGE="ghcr.io/$(echo "ROMANSHKVOLKOV" | tr '[:upper:]' '[:lower:]')/$(echo "ROMANSHKVOLKOV" | tr '[:upper:]' '[:lower:]'):$SHORT_SHA" \
  CADDY_HOST="domain.example.com" \
  CADDY_TLS="jose@guz-studio.dev" \
  DEPLOY_DB_DSN_MYSQL_ELEVA_CONTABO="db.dsn.eleva.contabo" \
  DEPLOY_DB_DSN_MYSQL_ELEVA="db.dsn.eleva" \
  DEPLOY_JWT_SECRET="jwt.secret" \
  DEPLOY_MAILER_EMAIL="mailer.email" \
  DEPLOY_MAILER_PASSWORD="mailer.password" \
  DEPLOY_MAILER_SMTP_HOST="mailer.smtp.host" \
  DEPLOY_MAILER_SMTP_PORT="mailer.smtp.port" \
  >./romanshkvolkov.prod.deployment.yml
