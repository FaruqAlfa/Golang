/*
AUTHENTICATION & AUTHORIZATION
Implementing Role-Based Access Control (RBAC) in Gin

Instructions: Create an HTTP server using the Gin framework that implements Role-Based Access Control (RBAC) to manage user access to different routes based on their role.

Instructions

Define User Roles: Users are assigned roles of admin, editor, or viewer.
admin can access /admin, /edit, and /view.
editor can access /edit and /view.
viewer can only access /view.
Implement Middleware: Write middleware in Gin that checks the user's role and permits access to routes accordingly.
Create Endpoints: Set up endpoints for /admin, /edit, and /view that use the middleware to restrict access based on the user's role.



HTML & CSS
Create this title & table in HTML & CSS


                                  <YOUR NAME: TABLE>

+----------------------------------+---------+------------------------+----------------+
|               Col1               |  Col2   |          Col3          | Numeric Column |
+----------------------------------+---------+------------------------+----------------+
| Value 1                          | Value 2 | 123                    |           10.0 |
| Separate                         | cols    | with a tab or 4 spaces |       -2,027.1 |
| This is a row with only one cell |         |                        |                |P
+----------------------------------+---------+------------------------+----------------+
Specifications
- <YOUR NAME: TABLE> is centered
- table header use blue color background, table body doesnt have to 
- table header is centered in each column
- table body use justify left