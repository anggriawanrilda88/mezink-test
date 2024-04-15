CREATE TABLE records ( 
	"id" Serial NOT NULL,
	"name" Character Varying( 50 ) NOT NULL,
	"marks" integer[] NOT NULL,
	"created_at" Timestamp With Time Zone DEFAULT now() NOT NULL,
	PRIMARY KEY ( "id" ) );
 ;

 CREATE INDEX "index_records_created_at" ON "records" USING btree( "created_at" );