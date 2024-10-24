# ************************************************************
# Antares - SQL Client
# Version 0.7.25
# 
# https://antares-sql.app/
# https://github.com/antares-sql/antares
# 
# Host: 127.0.0.1 (MySQL Community Server - GPL 8.4.2)
# Database: reddit_clone
# Generation time: 2024-10-24T23:02:49+02:00
# ************************************************************


/*!40101 SET @OLD_CHARACTER_SET_CLIENT=@@CHARACTER_SET_CLIENT */;
/*!40101 SET @OLD_CHARACTER_SET_RESULTS=@@CHARACTER_SET_RESULTS */;
/*!40101 SET @OLD_COLLATION_CONNECTION=@@COLLATION_CONNECTION */;
SET NAMES utf8mb4;
/*!40014 SET @OLD_FOREIGN_KEY_CHECKS=@@FOREIGN_KEY_CHECKS, FOREIGN_KEY_CHECKS=0 */;
/*!40101 SET @OLD_SQL_MODE=@@SQL_MODE, SQL_MODE='NO_AUTO_VALUE_ON_ZERO' */;
/*!40111 SET @OLD_SQL_NOTES=@@SQL_NOTES, SQL_NOTES=0 */;


# Dump of table comments
# ------------------------------------------------------------

LOCK TABLES `comments` WRITE;
/*!40000 ALTER TABLE `comments` DISABLE KEYS */;

INSERT INTO `comments` (`id_comment`, `comment`, `id_post`, `parent_id`, `id_user`) VALUES
	(1, "Fuck me that is so fucking unfunny...", 3, NULL, 3),
	(2, "Fuck you.", 3, 1, 4),
	(3, "That is so sweet to hear!", 2, NULL, 2);

/*!40000 ALTER TABLE `comments` ENABLE KEYS */;
UNLOCK TABLES;



# Dump of table posts
# ------------------------------------------------------------

LOCK TABLES `posts` WRITE;
/*!40000 ALTER TABLE `posts` DISABLE KEYS */;

INSERT INTO `posts` (`id_post`, `title`, `link`, `text`, `id_subreddit`, `id_user`) VALUES
	(1, "Phoronix: Several Linux Kernel Driver Maintainers Removed Due To Their Association To Russia", "https://www.phoronix.com/news/Russian-Linux-Maintainers-Drop", NULL, 1, 1),
	(2, "Finally did it, got my parent on linux.", NULL, "They’re in their 60’s, finally convinced them.  They say things like “This is the same…”  and I’m like  “Ya because that’s Firefox, the only program you use…”  “What was Windows even doing for us?”", 1, 3),
	(3, "What kind of crime is cockfighting?", NULL, "It\'s a feather-al offense.", 2, 4);

/*!40000 ALTER TABLE `posts` ENABLE KEYS */;
UNLOCK TABLES;



# Dump of table scores
# ------------------------------------------------------------

LOCK TABLES `scores` WRITE;
/*!40000 ALTER TABLE `scores` DISABLE KEYS */;

INSERT INTO `scores` (`id_score`, `score`, `id_post`, `id_comment`, `id_user`) VALUES
	(1, 0, NULL, 1, 1),
	(2, 0, NULL, 1, 5),
	(3, 1, NULL, 2, 1),
	(4, 1, NULL, 2, 5),
	(5, 1, 2, NULL, 2);

/*!40000 ALTER TABLE `scores` ENABLE KEYS */;
UNLOCK TABLES;



# Dump of table subreddits
# ------------------------------------------------------------

LOCK TABLES `subreddits` WRITE;
/*!40000 ALTER TABLE `subreddits` DISABLE KEYS */;

INSERT INTO `subreddits` (`id_subreddit`, `url`, `name`, `description`, `id_user`) VALUES
	(1, "linux", "Linux", "All things Linux.", 5),
	(2, "funny", "Funny", "If I see you posting cringe you will be banned.", 2);

/*!40000 ALTER TABLE `subreddits` ENABLE KEYS */;
UNLOCK TABLES;



# Dump of table users
# ------------------------------------------------------------

LOCK TABLES `users` WRITE;
/*!40000 ALTER TABLE `users` DISABLE KEYS */;

INSERT INTO `users` (`id_user`, `user_name`, `email`, `hashed_password`) VALUES
	(1, "Brice97", "Ayden.Wolff99@yahoo.com", "aayUXzZo8iD19ct"),
	(2, "Erna_Howe", "Jane56@hotmail.com", "cuA6jE_DYy3Gpoj"),
	(3, "Bertrand_Langosh48", "Devonte.Hilpert@gmail.com", "vY705g_vxicfOjw"),
	(4, "Rafaela_Cole88", "Naomie.Hilpert55@hotmail.com", "tOQLxSu3iFoNOL3"),
	(5, "Neil90", "Andres.Mraz43@gmail.com", "yhIV1rV3eNWcMAh");

/*!40000 ALTER TABLE `users` ENABLE KEYS */;
UNLOCK TABLES;



/*!40111 SET SQL_NOTES=@OLD_SQL_NOTES */;
/*!40101 SET SQL_MODE=@OLD_SQL_MODE */;
/*!40014 SET FOREIGN_KEY_CHECKS=@OLD_FOREIGN_KEY_CHECKS */;
/*!40101 SET CHARACTER_SET_CLIENT=@OLD_CHARACTER_SET_CLIENT */;
/*!40101 SET CHARACTER_SET_RESULTS=@OLD_CHARACTER_SET_RESULTS */;
/*!40101 SET COLLATION_CONNECTION=@OLD_COLLATION_CONNECTION */;

# Dump completed on 2024-10-24T23:02:49+02:00
