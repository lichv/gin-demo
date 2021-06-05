DROP TABLE IF EXISTS "public"."demo_whitelist_user";
CREATE TABLE "public"."demo_whitelist_user" (
  "code" varchar(64) COLLATE "pg_catalog"."default" NOT NULL DEFAULT ''::character varying,
  "data" text COLLATE "pg_catalog"."default" NOT NULL DEFAULT ''::character varying,
  "mtime" timestamptz(6) NOT NULL DEFAULT CURRENT_TIMESTAMP,
  "is_active" bool NOT NULL DEFAULT false,
  CONSTRAINT "demo_whitelist_user_pkey" PRIMARY KEY ("code")
)
;

ALTER TABLE "public"."demo_whitelist_user" OWNER TO "postgres";

COMMENT ON TABLE "public"."demo_whitelist_user" IS '用户白名单';

COMMENT ON COLUMN "public"."demo_whitelist_user"."code" IS '用户编号';

COMMENT ON COLUMN "public"."demo_whitelist_user"."data" IS '用户数据';

COMMENT ON COLUMN "public"."demo_whitelist_user"."mtime" IS '更新时间';

COMMENT ON COLUMN "public"."demo_whitelist_user"."is_active" IS '是否激活';
