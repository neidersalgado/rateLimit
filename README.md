**rateLimit - Deployment**

*Commands:*

1. **`make build`**
    - **Description:** Builds the Docker image for the rateLimit project.

2. **`make run`**
    - **Description:** Starts the Docker services in the background.

3. **`make down`**
    - **Description:** Stops the running Docker services.

4. **`make clean`**
    - **Description:** Cleans up Docker-related resources, removing unused images and containers.

*Deployment Instructions:*

To deploy the rateLimit project using Docker, follow these steps:

1. **Build Docker Image:**
    - Run `make build` to build the Docker image for the rateLimit project.

2. **Start Docker Services:**
    - Run `make run` to initiate the Docker services, allowing the project to run within a Dockerized environment.

3. **Stop Docker Services:**
    - To stop the Docker services, execute `make down`. This will halt the running containers.

4. **Clean Docker Resources:**
    - Execute `make clean` to clean up Docker-related resources, such as removing unused images and containers.

By following these commands, you can effectively deploy and manage the rateLimit project within a Docker environment.


**rateLimit**

*English:*

This project implements a rate limit to restrict the number of emails or notifications of a particular type sent to a user. The primary objective is to showcase programming skills, emphasizing the implementation of design patterns such as the Factory pattern and the Strategy pattern. Specifically, the illustrative use of the Strategy pattern demonstrates its utility for potential future implementations, allowing for different notification types through various mediums.

It's worth noting that while the current implementation of strategies appears similar, it has been designed for demonstrative purposes. The Strategy pattern provides the necessary structure for easy adaptation to future expansions and changes in notification strategies.

Additionally, unit testing has been integrated into key files to ensure the robustness and reliability of the code. The project employs a REST API to receive requests, containing information about the notification type, user, and message, showcasing the project's ability to efficiently handle incoming requests.

Concerning data persistence, an in-memory repository using maps has been implemented. However, due to the injectable implementation, this choice can be easily modified to suit different needs and environments, such as connecting to a database.

In summary, this project aims to provide a demonstration of good development practices, highlighting the use of design patterns, unit testing, and flexibility in implementation, showcasing the skills and approaches of the developer.

*Español:*

Este proyecto implementa un límite de velocidad para limitar el número de correos electrónicos o notificaciones de un tipo determinado que se envían a un usuario. El objetivo principal es exhibir habilidades de programación, destacando la implementación de patrones de diseño como el patrón Factory y el patrón Strategy. En particular, la utilización ilustrativa del patrón Strategy demuestra su utilidad para posibles implementaciones futuras, permitiendo diferentes tipos de notificaciones a través de diversos medios.

Es importante señalar que, aunque la implementación actual de las estrategias es similar, ha sido diseñada con fines demostrativos. El patrón Strategy proporciona la estructura necesaria para adaptarse fácilmente a futuras expansiones y cambios en las estrategias de notificación.

Además, se ha integrado pruebas unitarias en archivos clave para garantizar la robustez y confiabilidad del código. El proyecto utiliza una API REST para recibir solicitudes, que incluyen información sobre el tipo de notificación, el usuario y el mensaje, demostrando la capacidad del proyecto para gestionar y procesar eficientemente las peticiones entrantes.

En cuanto a la persistencia de datos, se ha implementado un repositorio en memoria utilizando maps. Sin embargo, gracias a la implementación inyectable, esta elección puede modificarse fácilmente para adaptarse a diferentes necesidades y entornos, como la conexión con una base de datos.

En resumen, este proyecto busca proporcionar una muestra de buenas prácticas de desarrollo, destacando el uso de patrones de diseño, pruebas unitarias y flexibilidad en la implementación, demostrando así las habilidades y enfoques del desarrollador.