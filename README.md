# CTF Challenge

Created for PCI Day at work. The challenge exploits basic web security principles and contains 2 flags.

## Premise

The website provides a shortlink service that allows you to generate shortened links that redirect from this website to another. The problem is that it has some major security holes in its design. 

## Solution

<details>
    <summary>Click here to view solution</summary>


### Flag 1: The admin token
- Player is intended to find that a few endpoints exist under `/admin` with endpoint fuzzing.
- When the user hits any of the `/admin` endpoints, it'll return that there's a header in the request for `X-Admin-Token`. 
- To find this token, they need to exploit that the `.env` file has been included in the build Docker image
- This can be accessed using the `/img/...` endpoint which allows for directory traversal

### Flag 2: the DB flag
- In the admin endpoint, the player will have access to an elevated db user string
- The `/admin` pages use unsanitised SQL for the query, so can perform SQL injection to find the `flags` table left in the DB, and extract the second flag

</details>