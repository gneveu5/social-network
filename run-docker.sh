cd backend
sudo docker build -t social-network-backend -f Dockerfile .
cd ../frontend
sudo docker build -t social-network-frontend .
sudo docker run -dp 8080:8080 social-network-backend
sudo docker run -p 3000:3000 social-network-frontend