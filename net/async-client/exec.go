package main

func FetchAll(urls []string, c *Client) {

	for _, url := range urls {
		go c.GetInAsync(url)
	}

}
