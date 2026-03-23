CREATE TABLE "service_lens_workload_lens" (
	"workloadId" uuid NOT NULL,
	"lensId" uuid NOT NULL
);
--> statement-breakpoint
CREATE TABLE "service_lens_workload_profile" (
	"workloadId" uuid NOT NULL,
	"profileId" uuid NOT NULL
);
--> statement-breakpoint
ALTER TABLE "service_lens_workload_lens" ADD CONSTRAINT "service_lens_workload_lens_workloadId_service_lens_workload_id_fk" FOREIGN KEY ("workloadId") REFERENCES "public"."service_lens_workload"("id") ON DELETE cascade ON UPDATE no action;--> statement-breakpoint
ALTER TABLE "service_lens_workload_lens" ADD CONSTRAINT "service_lens_workload_lens_lensId_service_lens_lens_id_fk" FOREIGN KEY ("lensId") REFERENCES "public"."service_lens_lens"("id") ON DELETE no action ON UPDATE no action;--> statement-breakpoint
ALTER TABLE "service_lens_workload_profile" ADD CONSTRAINT "service_lens_workload_profile_workloadId_service_lens_workload_id_fk" FOREIGN KEY ("workloadId") REFERENCES "public"."service_lens_workload"("id") ON DELETE cascade ON UPDATE no action;--> statement-breakpoint
ALTER TABLE "service_lens_workload_profile" ADD CONSTRAINT "service_lens_workload_profile_profileId_service_lens_profile_id_fk" FOREIGN KEY ("profileId") REFERENCES "public"."service_lens_profile"("id") ON DELETE no action ON UPDATE no action;