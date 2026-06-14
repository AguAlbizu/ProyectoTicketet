-- MySQL dump 10.13  Distrib 9.7.0, for macos15 (arm64)
--
-- Host: localhost    Database: ticketapp
-- ------------------------------------------------------
-- Server version	9.7.0

/*!40101 SET @OLD_CHARACTER_SET_CLIENT=@@CHARACTER_SET_CLIENT */;
/*!40101 SET @OLD_CHARACTER_SET_RESULTS=@@CHARACTER_SET_RESULTS */;
/*!40101 SET @OLD_COLLATION_CONNECTION=@@COLLATION_CONNECTION */;
/*!50503 SET NAMES utf8mb4 */;
/*!40103 SET @OLD_TIME_ZONE=@@TIME_ZONE */;
/*!40103 SET TIME_ZONE='+00:00' */;
/*!40014 SET @OLD_UNIQUE_CHECKS=@@UNIQUE_CHECKS, UNIQUE_CHECKS=0 */;
/*!40014 SET @OLD_FOREIGN_KEY_CHECKS=@@FOREIGN_KEY_CHECKS, FOREIGN_KEY_CHECKS=0 */;
/*!40101 SET @OLD_SQL_MODE=@@SQL_MODE, SQL_MODE='NO_AUTO_VALUE_ON_ZERO' */;
/*!40111 SET @OLD_SQL_NOTES=@@SQL_NOTES, SQL_NOTES=0 */;

--
-- Table structure for table `events`
--

DROP TABLE IF EXISTS `events`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE `events` (
  `id_events` bigint unsigned NOT NULL AUTO_INCREMENT,
  `titulo` varchar(200) COLLATE utf8mb4_unicode_ci NOT NULL,
  `descripcion` text COLLATE utf8mb4_unicode_ci,
  `fecha` datetime(3) NOT NULL,
  `hora` varchar(5) COLLATE utf8mb4_unicode_ci NOT NULL,
  `capacidad` bigint NOT NULL,
  `cupo_disponible` bigint NOT NULL,
  `categoria` varchar(100) COLLATE utf8mb4_unicode_ci DEFAULT NULL,
  `direccion` varchar(300) COLLATE utf8mb4_unicode_ci DEFAULT NULL,
  `imagen_url` varchar(500) COLLATE utf8mb4_unicode_ci DEFAULT NULL,
  `estado` varchar(20) COLLATE utf8mb4_unicode_ci NOT NULL DEFAULT 'activo',
  `created_at` datetime(3) DEFAULT NULL,
  `updated_at` datetime(3) DEFAULT NULL,
  `precio` int NOT NULL DEFAULT '0',
  PRIMARY KEY (`id_events`)
) ENGINE=InnoDB AUTO_INCREMENT=22 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `events`
--

LOCK TABLES `events` WRITE;
/*!40000 ALTER TABLE `events` DISABLE KEYS */;
INSERT INTO `events` VALUES (1,'Coldplay World Tour','El mejor show del año','2026-08-15 00:00:00.000','21:00',50000,50000,'Música','Estadio Mâs Monumental (River Plate) en Avenida Presidente Figueroa Alcorta 7597','https://img.asmedia.epimg.net/resizer/v2/474VCDG2YNAA3LA2FH52474IHE.jfif?auth=b970fec81a94b259e4a12f1f66b8b97da6393383b0a3d4d4b4817e754cd46be6&width=1472&height=828&focal=1000%2C794','activo',NULL,NULL,120000),(2,'Hamlet - Teatro Colón','Obra clásica de Shakespeare','2026-07-20 00:00:00.000','20:00',800,795,'Teatro','Teatro Colón en Buenos Aires es Cerrito 628, Ciudad Autónoma de Buenos Aires','https://www.artshub.co.uk/wp-content/uploads/sites/3/2025/03/Luke-Thallon-c-Marc-Brenner-1.jpg?w=1200','activo',NULL,NULL,35000),(3,'River vs Boca','Superclásico del fútbol argentino','2026-09-05 00:00:00.000','17:00',60000,59999,'Deportes','Estadio Mâs Monumental (River Plate) en Avenida Presidente Figueroa Alcorta 7597','https://prod-media.beinsports.com/image/River%20Plate%20vs%20Boca%20Juniors.png','activo',NULL,NULL,45000),(4,'Cosquín Rock 2026','La edición 26 del festival de rock más icónico de Argentina con +100 bandas: Abel Pintos, Franz Ferdinand, Lali Esposito, Eruca Sativa y más.','2026-02-14 00:00:00.000','14:00',110000,110000,'Música','Aeródromo Santa María de Punilla, Córdoba','https://encrypted-tbn0.gstatic.com/images?q=tbn:ANd9GcRycvHTIFknkkqGegxD5S9oWeONbH3noCAbzw&s','activo',NULL,NULL,55000),(5,'Lollapalooza Argentina 2026','11ª edición con Tyler The Creator, Sabrina Carpenter, Chappell Roan, Paulo Londra, Lorde, Skrillex y +100 artistas en 5 escenarios.','2026-03-13 00:00:00.000','12:30',300000,0,'Música','Hipódromo de San Isidro, Buenos Aires','https://yt3.googleusercontent.com/kNsdQfZZvifNR1ybRncn8YCBbW8Z-E8VcGJSIuFo0Rdi0ixb-3zm6iN_5GbkWNLxODeV-Dk1=s900-c-k-c0x00ffffff-no-rj','activo',NULL,NULL,90000),(6,'Fito Páez — Sale el Sol Tour','Gira nacional con versiones inéditas de sus grandes canciones y selección especial de su repertorio de cuatro décadas.','2026-03-19 00:00:00.000','21:00',15000,300,'Música','Movistar Arena, Humboldt 450, Villa Crespo, Buenos Aires','https://encrypted-tbn0.gstatic.com/images?q=tbn:ANd9GcTkiRh7TntSLPKbee8X9GkK1i3jbMAmy5cdhw&s','activo',NULL,NULL,55000),(7,'Talleres vs. Belgrano — Clásico Cordobés','El clásico más apasionado del interior del país enfrenta a Talleres y Belgrano en el estadio más grande de Córdoba. Ambiente electrizante con ambas hinchadas copando la ciudad.','2026-05-09 00:00:00.000','21:00',57000,0,'deporte','Estadio Mario Alberto Kempes, Bv. Juan Díaz de Solís s/n, Córdoba','https://www.365scores.com/es/news/wp-content/uploads/2026/05/PREVIA-93.jpg','activo','2026-06-13 23:31:43.000','2026-06-13 23:31:43.000',25000),(8,'Final Torneo Apertura 2026','River Plate vs. Belgrano definieron el campeón del Apertura 2026 en Córdoba.','2026-05-24 00:00:00.000','21:00',57000,0,'Deportes','Estadio Mario Alberto Kempes, Bv. Juan Díaz de Solís, Córdoba','https://inteligenciaargentina.ar/storage/1544/conversions/01KRW9N840YAMS4Z2D1GN0NVCM-hd.jpg','activo',NULL,NULL,35000),(9,'Argentina Premier Padel P1 — Buenos Aires','5ª edición del torneo P1 del circuito mundial Premier Padel en Argentina bajo techo.','2026-05-11 00:00:00.000','10:00',8000,8000,'Deportes','Estadio Mary Terán de Weiss, Parque Roca, Buenos Aires','https://encrypted-tbn0.gstatic.com/images?q=tbn:ANd9GcRPh8eH-TsS0Gl6U-nVxPjoPzkDsOPTjlcR8A&s','activo',NULL,NULL,45000),(10,'Ricardo Arjona — Lo que el Seco no dijo Tour','Residencia histórica con más de 17 fechas, considerada la producción más ambiciosa de su carrera.','2026-07-01 00:00:00.000','21:00',15000,150,'Música','Movistar Arena, Humboldt 450, Villa Crespo, Buenos Aires','https://encrypted-tbn0.gstatic.com/images?q=tbn:ANd9GcSMjxYkU1a8J-bBDkl_VaKY-e3niN4Z85jB1Q&s','activo',NULL,NULL,70000),(11,'Dua Lipa — Radical Optimism Tour','La estrella pop británica vuelve con su Radical Optimism Tour, agotando el Monumental en menos de una hora.','2026-11-07 00:00:00.000','21:00',84567,0,'Música','Estadio Monumental, Av. Figueroa Alcorta 7597, Núñez, Buenos Aires','https://upload.wikimedia.org/wikipedia/en/4/40/Radical_Optimism_Tour_poster.png','activo',NULL,NULL,130000),(12,'XIII Juegos Suramericanos Santa Fe 2026','Evento multideportivo con +4.000 atletas de 15 países en más de 57 disciplinas en tres ciudades sede.','2026-09-12 00:00:00.000','09:00',2000000,2000000,'Deportes','Rosario, Santa Fe y Rafaela, Provincia de Santa Fe','https://www.santafe.tur.ar/wp-content/uploads/sites/91/2026/01/cuenta-regresiva-juegos-suramericanos-2026-01-1024x576.jpg','activo',NULL,NULL,10000),(13,'Maratón Internacional de Buenos Aires — 42K','El maratón más rápido de Latinoamérica, abierto a corredores del mundo recorriendo las avenidas emblemáticas de CABA.','2026-09-20 00:00:00.000','07:00',15000,15000,'Deportes','Av. Figueroa Alcorta y Dorrego, Palermo, Buenos Aires','https://runtag.s3.amazonaws.com/original/m365238_logo-evento--22.png','activo',NULL,NULL,600),(14,'Media Maratón de Buenos Aires — 21K','La carrera de 21K más rápida de América, con categorías masculinas, femeninas y adaptadas.','2026-08-23 00:00:00.000','07:00',20000,20000,'Deportes','Av. Figueroa Alcorta y Dorrego, Palermo, Buenos Aires','https://encrypted-tbn0.gstatic.com/images?q=tbn:ANd9GcTU80vM4Eh3p-JzXDZVvmrftitwyS8QlGopnA&s','activo',NULL,NULL,3000),(15,'Arcángel — 20 Aniversario La 8va Maravilla Tour','El exponente del reggaetón clásico celebra dos décadas de su álbum icónico con producción a gran escala.','2026-08-28 00:00:00.000','21:00',15000,15000,'Música','Movistar Arena, Humboldt 450, Villa Crespo, Buenos Aires','https://encrypted-tbn0.gstatic.com/images?q=tbn:ANd9GcRwR7O65Vbh_3ikAHXOSkvcmzd3eBmAHNFLpg&s','activo',NULL,NULL,55000),(16,'Bad Bunny — Debí Tirar Más Fotos World Tour','Dos fechas agotadas en River con invitados especiales: Cazzu, Duki y Khea. Uno de los shows más comentados del año.','2026-02-13 00:00:00.000','21:00',84567,0,'Música','Estadio Monumental, Av. Figueroa Alcorta 7597, Núñez, Buenos Aires','https://encrypted-tbn0.gstatic.com/images?q=tbn:ANd9GcQSOYc9_VjJ4DmsJI-icmT4pbvNMoGji5DwOg&s','activo',NULL,NULL,140000),(17,'Festival de Teatro — FIBA Buenos Aires','Festival Internacional de Buenos Aires reuniendo compañías de teatro del mundo con funciones en distintos espacios.','2026-03-01 00:00:00.000','19:00',2000,2000,'Teatro','Varios teatros, Ciudad de Buenos Aires','https://www.diariodecultura.com.ar/wp-content/uploads/2023/02/FIBA-2023.jpg','activo',NULL,NULL,12000),(18,'Premio Clarín Espectáculos — Teatro','Ceremonia de premiación más importante del teatro argentino celebrando los mejores espectáculos de la temporada.','2026-12-01 00:00:00.000','20:00',1200,1200,'Teatro','Teatro Gran Rex, Av. Corrientes 857, Buenos Aires','https://encrypted-tbn0.gstatic.com/images?q=tbn:ANd9GcQAJ4X1hRdKVeWrcYIvRYUKBOBXg6BiyC99DA&s','activo',NULL,NULL,12000),(19,'Expo Agro 2026 — La Rural','La exposición agropecuaria más importante de América Latina con muestras ganaderas, tecnología y gastronomía.','2026-07-01 00:00:00.000','10:00',400000,400000,'Otro','Predio Ferial de Palermo, Av. Santa Fe 4363, Buenos Aires','https://www.expoagro.com.ar/wp-content/themes/expoagro-2022/assets/img/EA-20-logo.png','activo',NULL,NULL,12000),(20,'Carnaval de Buenos Aires 2026','Festejo popular con murgas, corsos barriales y bandas musicales en distintos barrios con entrada libre.','2026-02-01 00:00:00.000','20:00',35000,35000,'Otro','Distintos barrios de la Ciudad de Buenos Aires','https://fotos.perfil.com/2026/02/05/trim/987/555/carnaval-2026-05022026-2181238.jpg','activo',NULL,NULL,3000),(21,'Premier Padel Finals 2026','El cierre de la temporada 2026 con los 16 mejores jugadores del ranking FIP. Transmisión especial en Argentina vía Star+.','2026-12-15 00:00:00.000','18:00',10000,10000,'Deportes','Barcelona, España (Transmisión vía Star+)','https://encrypted-tbn0.gstatic.com/images?q=tbn:ANd9GcTkav_9_z7DPouAzWNLHKvLbpQ8kAcsNNHUVw&s','activo',NULL,NULL,70000);
/*!40000 ALTER TABLE `events` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `tickets`
--

DROP TABLE IF EXISTS `tickets`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE `tickets` (
  `id_tickets` bigint unsigned NOT NULL AUTO_INCREMENT,
  `id_users` bigint unsigned NOT NULL,
  `id_events` bigint unsigned NOT NULL,
  `estado` varchar(20) COLLATE utf8mb4_unicode_ci NOT NULL DEFAULT 'activo',
  `fecha_compra` datetime(3) NOT NULL,
  `created_at` datetime(3) DEFAULT NULL,
  `updated_at` datetime(3) DEFAULT NULL,
  `origen` varchar(20) COLLATE utf8mb4_unicode_ci NOT NULL DEFAULT 'compra',
  PRIMARY KEY (`id_tickets`),
  KEY `idx_tickets_id_users` (`id_users`),
  KEY `idx_tickets_id_events` (`id_events`)
) ENGINE=InnoDB AUTO_INCREMENT=21 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `tickets`
--

LOCK TABLES `tickets` WRITE;
/*!40000 ALTER TABLE `tickets` DISABLE KEYS */;
INSERT INTO `tickets` VALUES (1,1,1,'cancelado','2026-06-13 19:23:33.198','2026-06-13 19:23:33.198','2026-06-13 19:23:33.198','compra'),(2,1,1,'cancelado','2026-06-13 19:27:20.381','2026-06-13 19:27:20.381','2026-06-13 19:27:20.381','compra'),(3,2,1,'cancelado','2026-06-13 19:33:14.691','2026-06-13 19:33:14.691','2026-06-13 19:33:14.691','compra'),(4,1,1,'cancelado','2026-06-13 19:34:23.831','2026-06-13 19:34:23.831','2026-06-13 19:34:23.831','compra'),(5,1,2,'cancelado','2026-06-13 19:34:47.595','2026-06-13 19:34:47.595','2026-06-13 19:34:47.595','compra'),(6,1,1,'cancelado','2026-06-13 19:34:52.446','2026-06-13 19:34:52.446','2026-06-13 19:34:52.446','compra'),(7,1,1,'cancelado','2026-06-13 19:34:59.648','2026-06-13 19:34:59.649','2026-06-13 19:34:59.649','compra'),(8,3,1,'transferido','2026-06-13 19:54:20.653','2026-06-13 19:54:20.653','2026-06-13 19:54:20.653','compra'),(9,4,2,'transferido','2026-06-13 19:54:56.822','2026-06-13 19:54:56.822','2026-06-13 19:54:56.822','compra'),(10,3,2,'activo','2026-06-13 19:55:04.818','2026-06-13 19:55:04.819','2026-06-13 19:55:04.819','compra'),(11,3,3,'transferido','2026-06-13 19:55:35.021','2026-06-13 19:55:35.022','2026-06-13 19:55:35.022','compra'),(12,4,3,'transferido','2026-06-13 19:55:44.976','2026-06-13 19:55:44.976','2026-06-13 19:55:44.976','compra'),(13,4,2,'activo','2026-06-13 20:00:26.977','2026-06-13 20:00:26.977','2026-06-13 20:00:26.977','compra'),(14,3,3,'activo','2026-06-13 20:00:37.151','2026-06-13 20:00:37.151','2026-06-13 20:00:37.151','compra'),(15,2,2,'transferido','2026-06-13 20:01:29.621','2026-06-13 20:01:29.621','2026-06-13 20:01:29.621','compra'),(16,5,2,'activo','2026-06-13 20:01:29.643','2026-06-13 20:01:29.643','2026-06-13 20:01:29.643','transferencia'),(17,4,1,'activo','2026-06-13 20:02:15.514','2026-06-13 20:02:15.514','2026-06-13 20:02:15.514','transferencia'),(18,4,1,'activo','2026-06-13 22:10:48.232','2026-06-13 22:10:48.232','2026-06-13 22:10:48.232','compra'),(19,4,2,'activo','2026-06-13 23:48:06.556','2026-06-13 23:48:06.556','2026-06-13 23:48:06.556','compra'),(20,4,2,'activo','2026-06-13 23:48:12.986','2026-06-13 23:48:12.986','2026-06-13 23:48:12.986','compra');
/*!40000 ALTER TABLE `tickets` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `users`
--

DROP TABLE IF EXISTS `users`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE `users` (
  `id_users` bigint unsigned NOT NULL AUTO_INCREMENT,
  `nombre` varchar(100) COLLATE utf8mb4_unicode_ci NOT NULL,
  `email` varchar(150) COLLATE utf8mb4_unicode_ci NOT NULL,
  `password` varchar(255) COLLATE utf8mb4_unicode_ci NOT NULL,
  `rol` varchar(20) COLLATE utf8mb4_unicode_ci NOT NULL DEFAULT 'cliente',
  `created_at` datetime(3) DEFAULT NULL,
  `updated_at` datetime(3) DEFAULT NULL,
  PRIMARY KEY (`id_users`),
  UNIQUE KEY `idx_users_email` (`email`)
) ENGINE=InnoDB AUTO_INCREMENT=6 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `users`
--

LOCK TABLES `users` WRITE;
/*!40000 ALTER TABLE `users` DISABLE KEYS */;
INSERT INTO `users` VALUES (1,'Mario','12334@gmail.com','e68945648f0cac32cc96740003bb6c0b019707206bb2fda203e4b32e3c4925d3','cliente','2026-06-13 19:20:18.487','2026-06-13 19:20:18.487'),(2,'Test User','test@test.com','03ac674216f3e15c761ee1a5e255f067953623c8b388b4459e13f978d7c846f4','cliente','2026-06-13 19:33:02.033','2026-06-13 19:33:02.033'),(3,'Marian','truccom09@gmail.com','8d969eef6ecad3c29a3a629280e686cf0c3f5d5a86aff3ca12020c923adc6c92','cliente','2026-06-13 19:54:18.505','2026-06-13 19:54:18.505'),(4,'Mario','123@gmail.com','8d969eef6ecad3c29a3a629280e686cf0c3f5d5a86aff3ca12020c923adc6c92','cliente','2026-06-13 19:54:50.521','2026-06-13 19:54:50.521'),(5,'Receptor','receptor@test.com','03ac674216f3e15c761ee1a5e255f067953623c8b388b4459e13f978d7c846f4','cliente','2026-06-13 20:01:29.604','2026-06-13 20:01:29.604');
/*!40000 ALTER TABLE `users` ENABLE KEYS */;
UNLOCK TABLES;
/*!40103 SET TIME_ZONE=@OLD_TIME_ZONE */;

/*!40101 SET SQL_MODE=@OLD_SQL_MODE */;
/*!40014 SET FOREIGN_KEY_CHECKS=@OLD_FOREIGN_KEY_CHECKS */;
/*!40014 SET UNIQUE_CHECKS=@OLD_UNIQUE_CHECKS */;
/*!40101 SET CHARACTER_SET_CLIENT=@OLD_CHARACTER_SET_CLIENT */;
/*!40101 SET CHARACTER_SET_RESULTS=@OLD_CHARACTER_SET_RESULTS */;
/*!40101 SET COLLATION_CONNECTION=@OLD_COLLATION_CONNECTION */;
/*!40111 SET SQL_NOTES=@OLD_SQL_NOTES */;

-- Dump completed on 2026-06-14  0:03:29
