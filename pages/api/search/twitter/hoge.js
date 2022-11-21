import { Client, auth } from "twitter-api-sdk";

export default function twitter(req, res){
    // Initialize auth client first
    const authClient = new auth.OAuth2User({
        client_id: "YmVlSHNkTnppVVdnb3VkX1BjTmw6MTpjaQ",
        client_secret: "f0gnIs4LBJVxM5tyZMws1Nl2CSVKL7Qw_D_PMkI3yBORfz9fvB",
        callback: "YOUR-CALLBACK",
        scopes: ["tweet.read", "users.read", "offline.access"],
    });
   
    

    ( async () => {
        try {
            await authClient.requestAccessToken(code);
            console.log(authClient)
            // Pass auth credentials to the library client 
            const twitterClient = new Client(authClient);

            const fullArchiveSearch =
              await twitterClient.tweets.tweetsFullarchiveSearch({
                //One query/rule/filter for matching Tweets. Refer to https://t.co/rulelength to identify the max query length
                query: "(from:TwitterDev) new -is:retweet",
            });
            console.dir(fullArchiveSearch, {
              depth: null,
            });
          } catch (error) {
            console.log(error);
        }
    })();
}