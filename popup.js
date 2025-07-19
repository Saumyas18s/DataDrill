document.getElementById("fetch").addEventListener("click", async () => {
  const url = document.getElementById("url").value;
  if (!url) {
      alert("Please enter a URL");
      return;
  }

  try {
      const response = await fetch(`http://localhost:8080/fetch`, {
          method: "POST",
          headers: { "Content-Type": "application/json" },
          body: JSON.stringify({ url: url })
      });

      if (!response.ok) {
          throw new Error(`Failed to fetch website details: ${response.statusText}`);
      }

      const data = await response.json();
      console.log(data);

      const headingsList = document.getElementById("headingsList");
      headingsList.innerHTML = '';

      data.headings.forEach(heading => {
          const listItem = document.createElement("li");
          listItem.textContent = heading;
          headingsList.appendChild(listItem);
      });
  } catch (error) {
      console.error("Error fetching data:", error);
  }
});
