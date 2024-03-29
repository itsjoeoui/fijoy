ALTER TABLE "fijoy_account"
DROP CONSTRAINT "fijoy_account_workspace_id_fijoy_workspace_id_fk";

ALTER TABLE "fijoy_category"
DROP CONSTRAINT "fijoy_category_workspace_id_fijoy_workspace_id_fk";

ALTER TABLE "fijoy_tag"
DROP CONSTRAINT "fijoy_tag_workspace_id_fijoy_workspace_id_fk";

ALTER TABLE "fijoy_transaction"
DROP CONSTRAINT "fijoy_transaction_account_id_fijoy_account_id_fk";

ALTER TABLE "fijoy_transaction"
DROP CONSTRAINT "fijoy_transaction_category_id_fijoy_category_id_fk";

ALTER TABLE "fijoy_transaction"
DROP CONSTRAINT "workspace_user_reference";

ALTER TABLE "fijoy_workspace_user"
DROP CONSTRAINT "fijoy_workspace_user_workspace_id_fijoy_workspace_id_fk";

ALTER TABLE "fijoy_workspace_user"
DROP CONSTRAINT "fijoy_workspace_user_user_id_fijoy_user_id_fk";

ALTER TABLE "fijoy_user_key"
DROP CONSTRAINT "fijoy_user_key_user_id_fijoy_user_id_fk";
