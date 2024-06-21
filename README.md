Installatie

Voordat je deze package gebruikt, zorg ervoor dat je de benodigde dependencies installeert:
Installeer github.com/go-sql-driver/mysql met het commando go get -u github.com/go-sql-driver/mysql.
Configuratie

Gebruik JSON-configuratiebestanden om de databaseverbinding in te stellen:
cloudconfig.json: Configuratie voor de cloud database.
localconfig.json: Fallback configuratie voor een lokale database.

Gebruik
Verbind met de database en voer een kentekencontrole uit.

Functies
Connect()
Initialiseert de databaseverbinding met de opgegeven configuratiebestanden.

CheckKentekenInDatabase(kenteken string)
Controleert of een kenteken bestaat in de database.
Geeft eigenaarsgegevens terug als het kenteken bestaat.
