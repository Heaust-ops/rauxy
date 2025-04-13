# Rauxy

## Manage auth protected proxies for your local services.

Imagine you just built a state of the art PC.

I'm talking 2 RTX 5090s, Ryzen 9 9950X3D, 128GB of RAM, the works.

You have a local ollama server running at port 11434.

You're excited to port forward this so all your friends can use this to develop their apps too!

But there's a problem. ollama is unauthorized. if you forward port 11434, all of internet will have access to it, and your flagship PC!

---

Enter `rauxy`.

You can add auth tokens specifically for your friends,

`rauxy add ollama_rodrico 11434`

`rauxy add ollama_alex 11434`

List out those tokens,

`rauxy ls`

```
Output:
ID         Name                 Token                                    Port       Created At
1          ollama_rodrico       BNW2soyMmX6kQKMP36_Ax-7UOx1DqfXvQ_e0zWKM6nzGAwX6yfIWBydpWIL3GrSmFQnqGhSfJorDsX7Q-QFEfw 11434      2025-04-13 18:39:10.732936547+05:30
2          ollama_alex          YN-YMZMEtaI3EDSQVYP88CD5leRitafd0duo5rJg9biHk24jCmVQKpTD1E_phkYn-o9_-mlTNIMCyxWdzSxXCQ 11434      2025-04-13 18:39:14.668786975+05:30
```

Copy those tokens and send it over to your friends!

Now start a proxy on whatever port you wanna forward.

`rauxy serve 4123 11434`

And it's done!

Your friends can now use ollama on port `4123`! (given you forward it)

They just have to also pass their tokens as an `http` header, `"Authorization: Bearer <token>"`

otherwise it's a 1:1 proxy to ollama :)

Now say Rodico liked it so much, he built his own flagship PC, so he doesn't really need your api now.

You can revoke his token by `rauxy rm ollama_rodrico`

Easy as that!

---

You can build it yourself with `go build`

or use the pre-built binary, the file named `rauxy`

once you clone the repo, `chmod +x ./rauxy` to make the binary executable

and add a bash alias, `alias rauxy="/path/to/rauxy"` for easy global access!

Have fun! :)
