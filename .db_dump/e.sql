# ************************************************************
# Sequel Pro SQL dump
# Version 4541
#
# http://www.sequelpro.com/
# https://github.com/sequelpro/sequelpro
#
# Host: 127.0.0.1 (MySQL 5.7.21)
# Database: fairfax
# Generation Time: 2018-07-30 08:46:08 +0000
# ************************************************************


/*!40101 SET @OLD_CHARACTER_SET_CLIENT=@@CHARACTER_SET_CLIENT */;
/*!40101 SET @OLD_CHARACTER_SET_RESULTS=@@CHARACTER_SET_RESULTS */;
/*!40101 SET @OLD_COLLATION_CONNECTION=@@COLLATION_CONNECTION */;
/*!40101 SET NAMES utf8 */;
/*!40014 SET @OLD_FOREIGN_KEY_CHECKS=@@FOREIGN_KEY_CHECKS, FOREIGN_KEY_CHECKS=0 */;
/*!40101 SET @OLD_SQL_MODE=@@SQL_MODE, SQL_MODE='NO_AUTO_VALUE_ON_ZERO' */;
/*!40111 SET @OLD_SQL_NOTES=@@SQL_NOTES, SQL_NOTES=0 */;


# Dump of table articles
# ------------------------------------------------------------

DROP TABLE IF EXISTS `articles`;

CREATE TABLE `articles` (
  `id` int(11) unsigned NOT NULL AUTO_INCREMENT,
  `title` varchar(255) NOT NULL DEFAULT '',
  `body` longtext,
  `date` date NOT NULL,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8;

LOCK TABLES `articles` WRITE;
/*!40000 ALTER TABLE `articles` DISABLE KEYS */;

INSERT INTO `articles` (`id`, `title`, `body`, `date`)
VALUES
	(1,'Article 1','Lorem ipsum dolor amet fingerstache offal kinfolk, synth hot chicken poutine enamel pin tousled echo park farm-to-table chicharrones. Chicharrones prism godard glossier vaporware. Tote bag PBR&B ennui sriracha ugh pinterest viral truffaut pop-up bicycle rights pickled church-key. Kogi farm-to-table tbh distillery meditation lyft flannel, brunch 3 wolf moon subway tile salvia pabst dreamcatcher. Cold-pressed four loko lomo, kogi etsy before they sold out 8-bit shabby chic microdosing leggings. VHS sartorial squid vexillologist typewriter, fanny pack portland chicharrones marfa lomo hashtag neutra shaman. Fixie selfies slow-carb organic literally adaptogen, microdosing disrupt truffaut.\n\nAffogato bespoke venmo seitan vice franzen. Quinoa beard subway tile 90\'s, craft beer fanny pack church-key banh mi. Typewriter live-edge deep v taxidermy, banjo mustache shoreditch woke twee fingerstache. Quinoa microdosing ramps, aesthetic twee edison bulb helvetica slow-carb plaid fanny pack vegan. Disrupt chicharrones 90\'s wolf humblebrag man braid, keffiyeh trust fund fixie hashtag gochujang meggings lumbersexual meditation mustache.\n\nCornhole tattooed XOXO wayfarers taiyaki gentrify, hoodie cray salvia af pitchfork. 8-bit blue bottle vexillologist, taiyaki wayfarers ennui gluten-free whatever mumblecore narwhal try-hard shoreditch before they sold out. Messenger bag everyday carry yr tumeric thundercats pork belly. Tofu pitchfork cardigan, keffiyeh ugh squid semiotics green juice disrupt hexagon banh mi glossier taxidermy ethical authentic. Pork belly roof party ethical, semiotics hammock try-hard vape.\n\nDrinking vinegar knausgaard crucifix, four dollar toast umami succulents migas gochujang vaporware stumptown freegan tacos vinyl vape poke. Health goth gastropub photo booth small batch fixie literally, messenger bag lo-fi prism meditation ugh jianbing. Stumptown literally raw denim cardigan pabst, four dollar toast aesthetic kogi pork belly jianbing pinterest. Next level portland authentic organic letterpress. Fanny pack blog snackwave, letterpress deep v schlitz tote bag mlkshk PBR&B fam organic beard. Adaptogen portland hell of fanny pack kitsch pork belly yuccie literally palo santo you probably haven\'t heard of them gentrify helvetica brunch. Iceland tacos kale chips, distillery raclette mustache blue bottle aesthetic tofu vinyl banh mi cray.','2018-09-25'),
	(2,'Buzzwords 1','So we can hit the ground running, we will be iteratively offshoring every alignment in our space. Key players will take ownership of their ballpark figures by virtually integrating mobile organic growths. We aim to globally leverage our industry leader by conservatively relaying our corporate wholesale prince2 practitioners. Our Customer-Focused Brand solution offers enterprises a suite of company-wide offerings. Key players will take ownership of their dot-bombs by effectively transforming immersive clouds. In the milestone space, industry is strategically connecting its mission critical team players. So we can hit the ground running, we will be intelligently deep-diving every architecture in our space. You need to reliably engineer your verticals to increase your silo velocity. In the future, will you be able to ethically incentivize action points in your business? We use our holistic visibilities to dynamically manage our bandwidth expectations. Our business synergises deliverables to proactively and globally align our senior driver. Our standpoint development lifecycle enables competitive, innovative stacks. World-class paradigm shifts are becoming actionable proposition experts. We thrive because of our knowledge transfer user experience and end-to-end low hanging fruit culture. In the future, will you be able to reliably right-size standard setters in your business? Conservatively impacting effectively proactive market foci is crucial to our next-generation stakeholder.','2018-09-25'),
	(3,'Buzzwords 2','So we can hit the ground running, we will be iteratively offshoring every alignment in our space. Key players will take ownership of their ballpark figures by virtually integrating mobile organic growths. We aim to globally leverage our industry leader by conservatively relaying our corporate wholesale prince2 practitioners. Our Customer-Focused Brand solution offers enterprises a suite of company-wide offerings. Key players will take ownership of their dot-bombs by effectively transforming immersive clouds. In the milestone space, industry is strategically connecting its mission critical team players. So we can hit the ground running, we will be intelligently deep-diving every architecture in our space. You need to reliably engineer your verticals to increase your silo velocity. In the future, will you be able to ethically incentivize action points in your business? We use our holistic visibilities to dynamically manage our bandwidth expectations. Our business synergises deliverables to proactively and globally align our senior driver. Our standpoint development lifecycle enables competitive, innovative stacks. World-class paradigm shifts are becoming actionable proposition experts. We thrive because of our knowledge transfer user experience and end-to-end low hanging fruit culture. In the future, will you be able to reliably right-size standard setters in your business? Conservatively impacting effectively proactive market foci is crucial to our next-generation stakeholder.','2018-07-30'),
	(4,'Buzzwords 3','So we can hit the ground running, we will be iteratively offshoring every alignment in our space. Key players will take ownership of their ballpark figures by virtually integrating mobile organic growths. We aim to globally leverage our industry leader by conservatively relaying our corporate wholesale prince2 practitioners. Our Customer-Focused Brand solution offers enterprises a suite of company-wide offerings. Key players will take ownership of their dot-bombs by effectively transforming immersive clouds. In the milestone space, industry is strategically connecting its mission critical team players. So we can hit the ground running, we will be intelligently deep-diving every architecture in our space. You need to reliably engineer your verticals to increase your silo velocity. In the future, will you be able to ethically incentivize action points in your business? We use our holistic visibilities to dynamically manage our bandwidth expectations. Our business synergises deliverables to proactively and globally align our senior driver. Our standpoint development lifecycle enables competitive, innovative stacks. World-class paradigm shifts are becoming actionable proposition experts. We thrive because of our knowledge transfer user experience and end-to-end low hanging fruit culture. In the future, will you be able to reliably right-size standard setters in your business? Conservatively impacting effectively proactive market foci is crucial to our next-generation stakeholder.','2018-07-30'),
	(5,'Buzzwords 4','So we can hit the ground running, we will be iteratively offshoring every alignment in our space. Key players will take ownership of their ballpark figures by virtually integrating mobile organic growths. We aim to globally leverage our industry leader by conservatively relaying our corporate wholesale prince2 practitioners. Our Customer-Focused Brand solution offers enterprises a suite of company-wide offerings. Key players will take ownership of their dot-bombs by effectively transforming immersive clouds. In the milestone space, industry is strategically connecting its mission critical team players. So we can hit the ground running, we will be intelligently deep-diving every architecture in our space. You need to reliably engineer your verticals to increase your silo velocity. In the future, will you be able to ethically incentivize action points in your business? We use our holistic visibilities to dynamically manage our bandwidth expectations. Our business synergises deliverables to proactively and globally align our senior driver. Our standpoint development lifecycle enables competitive, innovative stacks. World-class paradigm shifts are becoming actionable proposition experts. We thrive because of our knowledge transfer user experience and end-to-end low hanging fruit culture. In the future, will you be able to reliably right-size standard setters in your business? Conservatively impacting effectively proactive market foci is crucial to our next-generation stakeholder.','2018-07-30'),
	(6,'Buzzwords 5','So we can hit the ground running, we will be iteratively offshoring every alignment in our space. Key players will take ownership of their ballpark figures by virtually integrating mobile organic growths. We aim to globally leverage our industry leader by conservatively relaying our corporate wholesale prince2 practitioners. Our Customer-Focused Brand solution offers enterprises a suite of company-wide offerings. Key players will take ownership of their dot-bombs by effectively transforming immersive clouds. In the milestone space, industry is strategically connecting its mission critical team players. So we can hit the ground running, we will be intelligently deep-diving every architecture in our space. You need to reliably engineer your verticals to increase your silo velocity. In the future, will you be able to ethically incentivize action points in your business? We use our holistic visibilities to dynamically manage our bandwidth expectations. Our business synergises deliverables to proactively and globally align our senior driver. Our standpoint development lifecycle enables competitive, innovative stacks. World-class paradigm shifts are becoming actionable proposition experts. We thrive because of our knowledge transfer user experience and end-to-end low hanging fruit culture. In the future, will you be able to reliably right-size standard setters in your business? Conservatively impacting effectively proactive market foci is crucial to our next-generation stakeholder.','2018-07-30'),
	(7,'Buzzwords 6','So we can hit the ground running, we will be iteratively offshoring every alignment in our space. Key players will take ownership of their ballpark figures by virtually integrating mobile organic growths. We aim to globally leverage our industry leader by conservatively relaying our corporate wholesale prince2 practitioners. Our Customer-Focused Brand solution offers enterprises a suite of company-wide offerings. Key players will take ownership of their dot-bombs by effectively transforming immersive clouds. In the milestone space, industry is strategically connecting its mission critical team players. So we can hit the ground running, we will be intelligently deep-diving every architecture in our space. You need to reliably engineer your verticals to increase your silo velocity. In the future, will you be able to ethically incentivize action points in your business? We use our holistic visibilities to dynamically manage our bandwidth expectations. Our business synergises deliverables to proactively and globally align our senior driver. Our standpoint development lifecycle enables competitive, innovative stacks. World-class paradigm shifts are becoming actionable proposition experts. We thrive because of our knowledge transfer user experience and end-to-end low hanging fruit culture. In the future, will you be able to reliably right-size standard setters in your business? Conservatively impacting effectively proactive market foci is crucial to our next-generation stakeholder.','2018-07-30'),
	(8,'Buzzwords 7','So we can hit the ground running, we will be iteratively offshoring every alignment in our space. Key players will take ownership of their ballpark figures by virtually integrating mobile organic growths. We aim to globally leverage our industry leader by conservatively relaying our corporate wholesale prince2 practitioners. Our Customer-Focused Brand solution offers enterprises a suite of company-wide offerings. Key players will take ownership of their dot-bombs by effectively transforming immersive clouds. In the milestone space, industry is strategically connecting its mission critical team players. So we can hit the ground running, we will be intelligently deep-diving every architecture in our space. You need to reliably engineer your verticals to increase your silo velocity. In the future, will you be able to ethically incentivize action points in your business? We use our holistic visibilities to dynamically manage our bandwidth expectations. Our business synergises deliverables to proactively and globally align our senior driver. Our standpoint development lifecycle enables competitive, innovative stacks. World-class paradigm shifts are becoming actionable proposition experts. We thrive because of our knowledge transfer user experience and end-to-end low hanging fruit culture. In the future, will you be able to reliably right-size standard setters in your business? Conservatively impacting effectively proactive market foci is crucial to our next-generation stakeholder.','2018-09-25'),
	(9,'Buzzwords 8','So we can hit the ground running, we will be iteratively offshoring every alignment in our space. Key players will take ownership of their ballpark figures by virtually integrating mobile organic growths. We aim to globally leverage our industry leader by conservatively relaying our corporate wholesale prince2 practitioners. Our Customer-Focused Brand solution offers enterprises a suite of company-wide offerings. Key players will take ownership of their dot-bombs by effectively transforming immersive clouds. In the milestone space, industry is strategically connecting its mission critical team players. So we can hit the ground running, we will be intelligently deep-diving every architecture in our space. You need to reliably engineer your verticals to increase your silo velocity. In the future, will you be able to ethically incentivize action points in your business? We use our holistic visibilities to dynamically manage our bandwidth expectations. Our business synergises deliverables to proactively and globally align our senior driver. Our standpoint development lifecycle enables competitive, innovative stacks. World-class paradigm shifts are becoming actionable proposition experts. We thrive because of our knowledge transfer user experience and end-to-end low hanging fruit culture. In the future, will you be able to reliably right-size standard setters in your business? Conservatively impacting effectively proactive market foci is crucial to our next-generation stakeholder.','2018-07-30'),
	(10,'Buzzwords 9','So we can hit the ground running, we will be iteratively offshoring every alignment in our space. Key players will take ownership of their ballpark figures by virtually integrating mobile organic growths. We aim to globally leverage our industry leader by conservatively relaying our corporate wholesale prince2 practitioners. Our Customer-Focused Brand solution offers enterprises a suite of company-wide offerings. Key players will take ownership of their dot-bombs by effectively transforming immersive clouds. In the milestone space, industry is strategically connecting its mission critical team players. So we can hit the ground running, we will be intelligently deep-diving every architecture in our space. You need to reliably engineer your verticals to increase your silo velocity. In the future, will you be able to ethically incentivize action points in your business? We use our holistic visibilities to dynamically manage our bandwidth expectations. Our business synergises deliverables to proactively and globally align our senior driver. Our standpoint development lifecycle enables competitive, innovative stacks. World-class paradigm shifts are becoming actionable proposition experts. We thrive because of our knowledge transfer user experience and end-to-end low hanging fruit culture. In the future, will you be able to reliably right-size standard setters in your business? Conservatively impacting effectively proactive market foci is crucial to our next-generation stakeholder.','2018-07-30'),
	(11,'Buzzwords 10','So we can hit the ground running, we will be iteratively offshoring every alignment in our space. Key players will take ownership of their ballpark figures by virtually integrating mobile organic growths. We aim to globally leverage our industry leader by conservatively relaying our corporate wholesale prince2 practitioners. Our Customer-Focused Brand solution offers enterprises a suite of company-wide offerings. Key players will take ownership of their dot-bombs by effectively transforming immersive clouds. In the milestone space, industry is strategically connecting its mission critical team players. So we can hit the ground running, we will be intelligently deep-diving every architecture in our space. You need to reliably engineer your verticals to increase your silo velocity. In the future, will you be able to ethically incentivize action points in your business? We use our holistic visibilities to dynamically manage our bandwidth expectations. Our business synergises deliverables to proactively and globally align our senior driver. Our standpoint development lifecycle enables competitive, innovative stacks. World-class paradigm shifts are becoming actionable proposition experts. We thrive because of our knowledge transfer user experience and end-to-end low hanging fruit culture. In the future, will you be able to reliably right-size standard setters in your business? Conservatively impacting effectively proactive market foci is crucial to our next-generation stakeholder.','2018-07-30');

/*!40000 ALTER TABLE `articles` ENABLE KEYS */;
UNLOCK TABLES;


# Dump of table tags
# ------------------------------------------------------------

DROP TABLE IF EXISTS `tags`;

CREATE TABLE `tags` (
  `id` int(11) unsigned NOT NULL AUTO_INCREMENT,
  `title` varchar(255) NOT NULL DEFAULT '',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8;

LOCK TABLES `tags` WRITE;
/*!40000 ALTER TABLE `tags` DISABLE KEYS */;

INSERT INTO `tags` (`id`, `title`)
VALUES
	(1,'science'),
	(2,'fitness'),
	(3,'health'),
	(4,'books'),
	(5,'synergy'),
	(6,'proactive'),
	(7,'industry'),
	(8,'enterprises'),
	(9,'space'),
	(10,'paradigm'),
	(11,'immersive'),
	(12,'generation'),
	(13,'organic'),
	(14,'stakeholder'),
	(15,'growths');

/*!40000 ALTER TABLE `tags` ENABLE KEYS */;
UNLOCK TABLES;


# Dump of table tags_for_articles
# ------------------------------------------------------------

DROP TABLE IF EXISTS `tags_for_articles`;

CREATE TABLE `tags_for_articles` (
  `id` int(11) unsigned NOT NULL AUTO_INCREMENT,
  `article_id` int(11) unsigned NOT NULL,
  `tag_id` int(11) unsigned NOT NULL,
  PRIMARY KEY (`id`),
  KEY `article_id` (`article_id`),
  KEY `tag` (`tag_id`),
  CONSTRAINT `article_id` FOREIGN KEY (`article_id`) REFERENCES `articles` (`id`),
  CONSTRAINT `tag` FOREIGN KEY (`tag_id`) REFERENCES `tags` (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8;

LOCK TABLES `tags_for_articles` WRITE;
/*!40000 ALTER TABLE `tags_for_articles` DISABLE KEYS */;

INSERT INTO `tags_for_articles` (`id`, `article_id`, `tag_id`)
VALUES
	(1,1,1),
	(2,1,2),
	(3,2,4),
	(4,2,5),
	(5,2,6),
	(6,3,7),
	(7,3,8),
	(8,3,6),
	(9,3,9),
	(10,4,10),
	(11,4,8),
	(12,4,6),
	(13,4,9),
	(14,5,10),
	(15,5,8),
	(16,5,6),
	(17,5,11),
	(18,6,10),
	(19,6,1),
	(20,6,6),
	(21,6,11),
	(22,7,12),
	(23,7,1),
	(24,7,6),
	(25,7,11),
	(26,8,12),
	(27,8,1),
	(28,8,6),
	(29,8,13),
	(30,9,14),
	(31,9,1),
	(32,9,6),
	(33,9,13),
	(34,10,14),
	(35,10,1),
	(36,10,15),
	(37,10,13),
	(38,11,7),
	(39,11,1),
	(40,11,15),
	(41,11,13);

/*!40000 ALTER TABLE `tags_for_articles` ENABLE KEYS */;
UNLOCK TABLES;



/*!40111 SET SQL_NOTES=@OLD_SQL_NOTES */;
/*!40101 SET SQL_MODE=@OLD_SQL_MODE */;
/*!40014 SET FOREIGN_KEY_CHECKS=@OLD_FOREIGN_KEY_CHECKS */;
/*!40101 SET CHARACTER_SET_CLIENT=@OLD_CHARACTER_SET_CLIENT */;
/*!40101 SET CHARACTER_SET_RESULTS=@OLD_CHARACTER_SET_RESULTS */;
/*!40101 SET COLLATION_CONNECTION=@OLD_COLLATION_CONNECTION */;
