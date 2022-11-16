from locust import HttpUser, task

class HelloWorldUser(HttpUser):
    @task
    def todos_test(self):
        self.client.get("http://localhost:8080/todos")