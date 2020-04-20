#!/bin/bash

# download bwa source (with task env info)
bget i bwa --verbose 2
# get all available keys
bget i -a
# in JSON format
bget i -a --format json
# view all bwa and samtools available tags in table
bget i bwa samtools -v
# view all bwa and samtools available tags in json
bget i bwa samtools -v --format json

# force download defuse reference (with task env info and save log to file)
bget i "reffa/defuse@GRCh38 #97" -t 10 -f --verbose 2 --save-log
bget i reffa/defuse@GRCh38 release=97 -t 10 -f
# download annovar reference
bget i db/annovar@clinvar_20170501 db/annovar@clinvar_20180603 builder=hg38

bget i db/annovar -v --formt text
bget i db/annovar version='clinvar_20131105, clinvar_20140211, clinvar_20140303, clinvar_20140702, clinvar_20140902, clinvar_20140929, clinvar_20150330, clinvar_20150629, clinvar_20151201, clinvar_20160302, clinvar_20161128, clinvar_20170130, clinvar_20170501, clinvar_20170905, clinvar_20180603, avsnp150, avsnp147, avsnp144, avsnp142, avsnp138, cadd, caddgt10, caddgt20, cadd13, cadd13gt10, cadd13gt20, cg69, cg46, cosmic70, cosmic68wgs, cosmic68, cosmic67wgs, cosmic67, cosmic65, cosmic64, dbnsfp35a, dbnsfp33a, dbnsfp31a_interpro, dbnsfp30a, dbscsnv11, eigen, esp6500siv2_ea, esp6500siv2_aa, esp6500siv2_all, exac03nontcga, exac03nonpsych, exac03, fathmm, gerp++gt2, gme, gnomad_exome, gnomad_genome, gwava, hrcr1, icgc21, intervar_20170202, kaviar_20150923, ljb26_all, mcap, mitimpact2, mitimpact24, nci60, popfreq_max_20150413, popfreq_all_20150413, revel, regsnpintron' builder=hg19 -t 10 -f

bget --clean

