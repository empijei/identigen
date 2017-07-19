package lists

// Addresses exports an address for a city where a Partita Iva can be created
var Addresses = map[string]string{
	"Torino":                         "Piazza Castello, 15, 10124 Torino TO, Italia",
	"Vercelli-Biella":                "Piazza Cesare Battisti, 6, 13100 Vercelli VC, Italia",
	"Novara-Verbania":                "Corso Felice Cavallotti, 1, 28100 Novara NO, Italia",
	"Cuneo":                          "Corso Nizza, 48, 12100 Cuneo CN, Italia",
	"Asti":                           "Corso Dante Alighieri, 5, 14100 Asti AT, Italia",
	"Alessandria":                    "Spalto Borgoglio, 2, 15121 Alessandria AL, Italia",
	"Aosta":                          "Viale Partigiani, 4, 11100 Aosta AO, Italia",
	"Imperia":                        "Via Ospedale, 31, 18100 Imperia IM, Italia",
	"Savona":                         "Via Cristoforo Astengo, 50, 17100 Savona SV, Italia",
	"Genova":                         "Piazza Giuseppe Verdi, 400, 16121 Genova GE, Italia",
	"La Spezia":                      "Viale Italia, 12, 19121 La Spezia SP, Italia",
	"Varese":                         "Via Cesare Battisti, 1, 21100 Varese VA, Italia",
	"Como-Lecco":                     "Via dei Mille, 20, 22100 Como CO, Italia",
	"Sondrio":                        "Via Pelosi, 3, 23100 Sondrio SO, Italia",
	"Milano-Lodi":                    "Via Dante, 1, 20121 Milano MI, Italia",
	"Bergamo":                        "Via G. Camozzi, 127, 24121 Bergamo BG, Italia",
	"Brescia":                        "Via dei Mille, 5, 25122 Brescia BS, Italia",
	"Pavia":                          "Corso Giuseppe Mazzini, 11, 27100 Pavia PV, Italia",
	"Cremona":                        "Via Brescia, 54c, 26100 Cremona CR, Italia",
	"Mantova":                        "Corso della Libertà, 17, 46100 Mantova MN, Italia",
	"Bolzano":                        "Via della Mostra, 8, 39100 Bolzano BZ, Italia",
	"Trento":                         "Via Giovanni Segantini, 18, 38121 Trento TN, Italia",
	"Verona":                         "Via Roma, 1, 37121 Verona VR, Italia",
	"Vicenza":                        "Viale Giuseppe Mazzini, 5, 36100 Vicenza VI, Italia",
	"Belluno":                        "Via Attilio Tissi, 29, 32100 Belluno BL, Italia",
	"Treviso":                        "Piazza Duomo, 3, 31100 Treviso TV, Italia",
	"Venezia":                        "Ponte della Libertà, 302A, 30135 Venezia VE, Italia",
	"Padova":                         "Via Daniele Manin, 1, 35122 Padova PD, Italia",
	"Rovigo":                         "Viale Regina Margherita, 41, 45100 Rovigo RO, Italia",
	"Udine":                          "Viale S. Daniele, 92, 33100 Udine UD, Italia",
	"Gorizia":                        "Viale XXIV Maggio, 20, 34170 Gorizia GO, Italia",
	"Trieste":                        "Corso Umberto Saba, 1, 34122 Trieste TS, Italia",
	"Piacenza":                       "Viale Palmerio Raimondo, 4, 29121 Piacenza PC, Italia",
	"Parma":                          "Piazza Giuseppe Garibaldi, 12, 43121 Parma PR, Italia",
	"Reggio Emilia":                  "Via Emilia Santo Stefano, 6, 42121 Reggio Emilia RE, Italia",
	"Modena":                         "Via Emilia, 207, 41121 Modena MO, Italia",
	"Bologna":                        "Via dell'Indipendenza, 1, 40125 Bologna BO, Italia",
	"Ferrara":                        "Largo Castello, 1, 44121 Ferrara FE, Italia",
	"Ravenna":                        "Via di Roma, 164, 48121 Ravenna RA, Italia",
	"Forlì-Rimini":                   "Piazza Saffi Aurelio, 8, 47121 Forlì FC, Italia",
	"Pesaro":                         "Largo Aldo Moro, 12, 61121 Pesaro PU, Italia",
	"Ancona":                         "Via Palestro, 20, 60122 Ancona AN, Italia",
	"Macerata":                       "Viale Trieste, 9, 62100 Macerata MC, Italia",
	"Ascoli Piceno":                  "Via Dino Angelini, 147, 63100 Ascoli Piceno AP, Italia",
	"Massa Carrara":                  "Via Saldina, 7, 54011 Saldina MS, Italia",
	"Lucca":                          "Via Roma, 5, 55100 Lucca LU, Italia",
	"Pistoia":                        "Via Sergio Sacconi, 19, 51100 Pistoia PT, Italia",
	"Firenze":                        "Via dei Calzaiuoli, 50122 Firenze FI, Italia",
	"Livorno":                        "Via dei Lanzi, 1, 57123 Livorno LI, Italia",
	"Pisa":                           "Via Contessa Matilde, 1, 56126 Pisa PI, Italia",
	"Arezzo":                         "Via Giuseppe Garibaldi, 94, 52100 Arezzo AR, Italia",
	"Siena":                          "Via di Città, 36, 53100 Siena SI, Italia",
	"Grosseto":                       "Piazza Fratelli Rosselli, 16, 58100 Grosseto GR, Italia",
	"Perugia":                        "Via Giuseppe Mazzini, 15, 06121 Perugia PG, Italia",
	"Terni":                          "Via Carducci, 7, 05100 Terni TR, Italia",
	"Viterbo":                        "Piazza Giuseppe Verdi, 2, 01100 Viterbo VT, Italia",
	"Rieti":                          "Viale Emilio Maraini, 02100 Rieti RI, Italia",
	"Roma":                           "Piazza della Repubblica, 10, 00185 Roma RM, Italia",
	"Latina":                         "Piazza del Popolo, 1, 04100 Latina LT, Italia",
	"Frosinone":                      "Via Aldo Moro, 243, 03100 Frosinone FR, Italia",
	"Caserta":                        "Piazza S. Sebastiano, 5, 81100 Caserta CE, Italia",
	"Benevento":                      "Viale dei Rettori, 9, 82100 Benevento BN, Italia",
	"Napoli":                         "Corso Giuseppe Garibaldi, 1115, 80142 Napoli NA, Italia",
	"Avellino":                       "Via Malta, 16, 83100 Avellino AV, Italia",
	"Salerno":                        "Corso Vittorio Emanuele, 230, 84100 Salerno SA, Italia",
	"L'Aquila":                       "Corso Vittorio Emanuele, 60, 67100 L'Aquila AQ, Italia",
	"Teramo":                         "Corso S. Giorgio, 4, 64100 Teramo TE, Italia",
	"Pescara":                        "Via Firenze, 6, 65126 Pescara PE, Italia",
	"Chieti":                         "Via Asinio Herio, 101, 66100 Chieti CH, Italia",
	"Campobasso":                     "Corso Giuseppe Mazzini, 33, 86100 Campobasso CB, Italia",
	"Foggia":                         "Corso Giuseppe Garibaldi, 31, 71121 Foggia FG, Italia",
	"Bari":                           "Via Giuseppe Capruzzi, 126, 70125 Bari BA, Italia",
	"Taranto":                        "Via Crispi, 104, 74123 Taranto TA, Italia",
	"Brindisi":                       "Corso Roma, 146, 72100 Brindisi BR, Italia",
	"Lecce":                          "Via G. Marconi, 35, 73100 Lecce LE, Italia",
	"Potenza":                        "Via Pretoria, 260, 85100 Potenza PZ, Italia",
	"Matera":                         "Via Don Giovanni Minzoni, 2, 75100 Matera MT, Italia",
	"Cosenza":                        "Via Calabria, 6, 87100 Cosenza CS, Italia",
	"Catanzaro-Crotone-ViboValentia": "Piazza G. Matteotti, 15, 88100 Catanzaro CZ, Italia",
	"Reggio Calabria":                "Via Palamolla, 8, 89125 Reggio Calabria RC, Italia",
	"Trapani":                        "Via delle Acacie, 5, 91100 Trapani TP, Italia",
	"Palermo":                        "Via Maqueda, 271, 90133 Palermo PA, Italia",
	"Messina":                        "Via Sant'Agostino, 2, 98122 Messina ME, Italia",
	"Agrigento":                      "Piazza Luigi Pirandello, 33, 92100 Agrigento AG, Italia",
	"Caltanissetta":                  "Corso Vittorio Emanuele II, 97, 93100 Caltanissetta CL, Italia",
	"Enna":                           "Via Mercato, 2, 94100 Enna EN, Italia",
	"Catania":                        "Piazza S. Domenico, 6, 95124 Catania CT, Italia",
	"Ragusa":                         "Via Roma, 180, 97100 Ragusa RG, Italia",
	"Siracusa":                       "Viale Teocrito, 74, 96100 Siracusa SR, Italia",
	"Sassari":                        "Piazza Santa Maria, 110, 07100 Sassari SS, Italia",
	"Nuoro":                          "Via la Marmora, 99, 08100 Nuoro NU, Italia",
	"Cagliari":                       "Via Sant'Alenixedda, 7, 09128 Cagliari CA, Italia",
	"Pordenone":                      "Largo S. Giovanni, 31, 33170 Pordenone PN, Italia",
	"Isernia":                        "Corso Giuseppe Garibaldi, 14, 86170 Isernia IS, Italia",
	"Oristano":                       "Via Cagliari, 111, 09170 Oristano OR, Italia",
	"Milano 2 (Monza)":               "Via Italia, 12, 20052 Monza MB, Italia",
	"Firenze 2 (Prato)":              "Via Garibaldi, 50, 59100 Prato PO, Italia",
	"Brescia 2":                      "Via Giovanni Battista Tiepolo 25124 Brescia BS, Italia",
	"Genova 2":                       "Via Coronata, 100 16152 Genova GE, Italia",
	"Roma 2":                         "Via delle Fornaci, 253 00165 Roma RM, Italia",
	"Bologna 2":                      "Via della Salute, 11 40132 Bologna BO, Italia",
	"Napoli 2":                       "Via Leonardo Bianchi, 23 80131 Napoli NA, Italia",
}
