# Online Bookstore Project

## Overview

The Online Bookstore Project is a simple online bookstore system developed using gRPC and Go. It allows users to browse books, add them to their cart, and place orders.

## Features

- **User Authentication:**
  - Sign up, log in, and manage user profiles.

- **Book Listing:**
  - Retrieve a list of available books with details such as title, author, genre, and price.

- **Book Details:**
  - View detailed information about a specific book.

- **Shopping Cart:**
  - Add books to the cart, view cart contents, and remove items.

- **Order Placement:**
  - Place orders, deduct items from the inventory, and update user order history.

- **Order History:**
  - View order history with details such as order date, ordered items, and total amount.

- **Inventory Management:**
  - Basic inventory system to track book availability.

- **Search Functionality:**
  - Find books based on keywords, authors, or genres.

- **Rating and Reviews (Optional):**
  - Leave ratings and reviews for books.

- **Concurrency Control:**
  - Handle concurrent updates to the inventory to prevent race conditions.

- **Error Handling:**
  - Robust error handling for scenarios like out-of-stock items or failed order placements.

- **User Dashboard:**
  - Manage profiles, view order history, and track user activity.
