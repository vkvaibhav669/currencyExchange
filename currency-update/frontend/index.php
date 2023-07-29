<!DOCTYPE html>
<html>
<head>
    <title>Exchange Rates</title>
</head>
<body>
    <h1>Exchange Rates</h1>

    <?php
    // Function to fetch exchange rates from the back end API
    function fetchExchangeRates() {
        $apiUrl = 'http://backend:8080/api/exchange-rates'; // Docker service name "backend" points to the Golang backend

        $ch = curl_init($apiUrl);
        curl_setopt($ch, CURLOPT_RETURNTRANSFER, true);

        $response = curl_exec($ch);
        $httpCode = curl_getinfo($ch, CURLINFO_HTTP_CODE);

        curl_close($ch);

        if ($httpCode === 200) {
            return json_decode($response, true);
        } else {
            return null;
        }
    }

    // Fetch exchange rates from the back end
    $exchangeRates = fetchExchangeRates();

    if ($exchangeRates) {
        echo '<h2>Exchange Rates for the Last 10 Days</h2>';
        echo '<table border="1">';
        echo '<tr><th>Currency</th><th>Rate</th><th>Date</th></tr>';

        foreach ($exchangeRates as $rate) {
            echo '<tr>';
            echo '<td>' . $rate['currency'] . '</td>';
            echo '<td>' . $rate['rate'] . '</td>';
            echo '<td>' . $rate['date'] . '</td>';
            echo '</tr>';
        }

        echo '</table>';
    } else {
        echo '<p>Failed to fetch exchange rates. Please try again later.</p>';
    }
    ?>
</body>
</html>
