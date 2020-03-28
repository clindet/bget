dat <- read.csv("doi.list.oa.journal.txt", sep = "\t")

ref <- openxlsx::read.xlsx("~/Documents/paper/bget/data/oa_allow_spider_with_doi_and_num_of_paper_plus100.xlsx")

final <- merge(dat, ref, by.x = 4, by.y=4, all.x = T)
colnames(final)[1] <- "Journal.ISSN.(print.version)"

final2 <-  merge(dat, ref, by.x = 4, by.y=5, all.x = T)
colnames(final2)[1] <- "Journal.EISSN.(online.version)"
final2 <- final2[,colnames(final)]
final <- rbind(final, final2)

final <- final[!is.na(final$Publisher),]

write.table(final, "doi.list.oa.journal2.txt", sep = "\t", row.names = F, quote = F)