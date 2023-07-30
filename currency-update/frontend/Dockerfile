# Use the official PHP image as the base image
FROM php:latest

# Set the working directory inside the container
WORKDIR /var/www/html

# Copy the PHP application files to the container's working directory
COPY index.php .

# Expose the port that the PHP application listens on (if needed)
EXPOSE 80

# Command to run a PHP server (you can replace "index.php" with your entry point file if different)
CMD ["php", "-S", "0.0.0.0:80", "index.php"]
