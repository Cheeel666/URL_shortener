sudo docker stop url_shortener_db
sudo docker rm url_shortener_db
sudo docker build -t url_shortener-back -f Dockerfile .
sudo docker run -p 5432:5432 --name url_shortener -t url_shortener-db
sudo docker run -p 8080:8080 --name url_shortener-back -t url_shortener-back