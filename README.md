# uca-edt

A simple tool to generate a schedule for the UCA's students.

## Environment

This project needs to include the following .env file:

```
MONGO_URL="mongodb://root:password@ip:port"
SECURITY_KEY="your_secret_key"
```

## Admin endpoint 

The admin endpoint is available at `/admin` and is protected by a basic auth. The secret key is defined in the .env file as `SECURITY_KEY`.
To access the admin endpoint, you need to provide the key as an HTTP header: `Key: your_secret_key`.

We might change this in the future to use a more secure authentication method (like JWT) but for now, this is the easiest way to do it, and it might be enough for our needs.
If you have any suggestion, feel free to open an issue or a PR.

## Use our project for your own school

If you want to use our project for your own school, you can contact us. We will be happy to help you. 
This project is open source, so you can also fork it and adapt it to your needs.

## Disclaimer

This project is not affiliated with the UCA in any way. It is a personal project that I made for my own use. I am not responsible for any damage caused by the use of this project.

## License

This project is licensed under the MIT License - see the [LICENSE](LICENSE) file for details

## Acknowledgments

* [UCA](https://univ-cotedazur.fr/) - The University of Nice Côte d'Azur
* [Triformine](https://github.com/TriForMine) - Edt parser & front-end
  * [Deployment](https://uca-edt.triformine.dev/) - The deployment of the project
* [UCA's science Discord](https://discord.gg/XYX5gNxPP4) - The community we are part of & developed this project for

## Authors

* **[CorentinGS](https://github.com/CorentinGS)**
* **[TriForMine](https://github.com/TriForMine)**