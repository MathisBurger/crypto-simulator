<div align="center">
    <h1>Crypto Simulator</h1>
    <hr>
    <strong>
    A multiplayer cryptocurrency simulation game for everyone.
</strong><br><br>
<img src="https://img.shields.io/github/languages/count/mathisburger/crypto-simulator?style=for-the-badge">
<img src="https://img.shields.io/github/languages/top/mathisburger/crypto-simulator?style=for-the-badge">
<img src="https://img.shields.io/tokei/lines/github/mathisburger/crypto-simulator?style=for-the-badge">
</div>

<hr>

<div align="center">
<img src="https://upload.wikimedia.org/wikipedia/commons/thumb/0/05/Go_Logo_Blue.svg/1200px-Go_Logo_Blue.svg.png" height="100">
<img src="https://upload.wikimedia.org/wikipedia/commons/thumb/c/cf/Angular_full_color_logo.svg/1200px-Angular_full_color_logo.svg.png" height="120">
<img src="https://coincap.io/static/icons/favicon.ico" height="100">
</div>


# Information

---
The crypto simulator is an open source simulation game based on an angular web application.
You can trade with virtual currencies, but with real coin prices. Try it out yourself by 
creating an account on <a href="https://crypto-simulator.mathis-burger.de">
the main instance</a>.<br>
A service is calling the <a href="https://api.coincap.io/v2/assets">coincap API</a> every 10 seconds, and saves the 
current prices and more information about the currency into the database.<br>
If a user wants to buy a specific amount of a currency, the backend checks if the user has got enough money to buy this.
If is able to buy the currency, an entry with information about the trade, including the current price, is beeing saved
into the database.


# Idea

---
I had the idea to create the crypto simulator, because I am a highschool student. 
Therefore I am not allowed to trade with crypto currencies. But I was really interested
into this topic and wanted to learn more about crypto. 
I created this game, to learn more about crypto currencies and learn how to trade them.