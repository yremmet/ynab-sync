# YNAB-Sync

This is a tool to sync your transactions from MoneyMoney to YNAB. 💶   
🚨It's only a proof of concept with no error handling, useful logs or any help whatsoever.🚧  
🚨It can and most certainly will mess up your transaction history. 🚧

Maybe it will receive more love and evolve to something more useful, but maybe not. 🤷‍♂️


##### If you want to try it anyways: 

1. go build 
2. create a config in `~/.ynab-sync/config.json` containing your personal access token for YNAB. [(YNAB API Documentation↗️)️](https://api.youneedabudget.com)
    ```
    {"PersonalAccessToken":"$ACCESS_TOKEN"}
    ```
3. Add the Bank Account Number or Credit Card Number to your YNAB accounts. 
4. Open and Unlock MoneyMoney.
5. Run the application 


