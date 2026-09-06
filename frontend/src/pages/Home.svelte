<script>
    let estimatedMonthlyIncome
    let currentMonthlyIncome
    let totalInvestments
    let totalDebt
    let periodSpending
    let periodSpendingByCategory
    let startMonth = 1
    let startYear = 2026
    let endMonth = 9
    let endYear = 2026
</script>

<div class="home-page">
    <h1>Home Dashboard</h1>
    <div class="dashboard">

        <label for="month">Start Date</label>
        <select name="month" id="month" bind:value={startMonth}>
            <option value={1}>January</option>
            <option value={2}>February</option>
            <option value={3}>March</option>
            <option value={4}>April</option>
            <option value={5}>May</option>
            <option value={6}>June</option>
            <option value={7}>July</option>
            <option value={8}>August</option>
            <option value={9}>September</option>
            <option value={10}>October</option>
            <option value={11}>November</option>
            <option value={12}>December</option>
        </select>

        <select name="year" id="year" bind:value={startYear}>
            <option value={2000}>2000</option>
            <option value={2001}>2001</option>
            <option value={2002}>2002</option>
            <option value={2003}>2003</option>
            <option value={2004}>2004</option>
            <option value={2005}>2005</option>
            <option value={2006}>2006</option>
            <option value={2007}>2007</option>
            <option value={2008}>2008</option>
            <option value={2009}>2009</option>
            <option value={2010}>2010</option>
            <option value={2011}>2011</option>
            <option value={2012}>2012</option>
            <option value={2013}>2013</option>
            <option value={2014}>2014</option>
            <option value={2015}>2015</option>
            <option value={2016}>2016</option>
            <option value={2017}>2017</option>
            <option value={2018}>2018</option>
            <option value={2019}>2019</option>
            <option value={2020}>2020</option>
            <option value={2021}>2021</option>
            <option value={2022}>2022</option>
            <option value={2023}>2023</option>
            <option value={2024}>2024</option>
            <option value={2025}>2025</option>
            <option value={2026}>2026</option>
        </select>

        <label for="monthEnd">End Date</label>
        <select name="monthEnd" id="monthEnd" bind:value={endMonth}>
            <option value={1}>January</option>
            <option value={2}>February</option>
            <option value={3}>March</option>
            <option value={4}>April</option>
            <option value={5}>May</option>
            <option value={6}>June</option>
            <option value={7}>July</option>
            <option value={8}>August</option>
            <option value={9}>September</option>
            <option value={10}>October</option>
            <option value={11}>November</option>
            <option value={12}>December</option>
        </select>

        <select name="yearEnd" id="yearEnd" bind:value={endYear}>
            <option value={2000}>2000</option>
            <option value={2001}>2001</option>
            <option value={2002}>2002</option>
            <option value={2003}>2003</option>
            <option value={2004}>2004</option>
            <option value={2005}>2005</option>
            <option value={2006}>2006</option>
            <option value={2007}>2007</option>
            <option value={2008}>2008</option>
            <option value={2009}>2009</option>
            <option value={2010}>2010</option>
            <option value={2011}>2011</option>
            <option value={2012}>2012</option>
            <option value={2013}>2013</option>
            <option value={2014}>2014</option>
            <option value={2015}>2015</option>
            <option value={2016}>2016</option>
            <option value={2017}>2017</option>
            <option value={2018}>2018</option>
            <option value={2019}>2019</option>
            <option value={2020}>2020</option>
            <option value={2021}>2021</option>
            <option value={2022}>2022</option>
            <option value={2023}>2023</option>
            <option value={2024}>2024</option>
            <option value={2025}>2025</option>
            <option value={2026}>2026</option>
        </select>

        <br>
        <h3>Line chart for Income bs Spending or Stacked Bar Chart Groups  </h3>
        <h3>Have Spending By Category Breakdown In List Below chart on left </h3>
        <h3>Have top 5 Non-Expense Purchases for the range blow chart on right </h3>
        <h3>Have top 5 Non-Expense Spend By Company under top purchases </h3>

    </div>
</div>

<style>
    .home-page {
        display: flex;
        flex-direction: column;
        align-items: center;
        border: #4a90e2 2px solid;
    }

    select {
        padding: 8px 12px;
        margin: 5px;
        font-size: 16px;
        font-family: Arial, sans-serif;
        color: #333;
        background-color: #fff;
        border: 1px solid #ccc;
        border-radius: 6px;
        cursor: pointer;
        outline: none;
        transition: border-color 0.2s ease, box-shadow 0.2s ease;
    }

    select:hover {
        border-color: #999;
    }

    select:focus {
        border-color: #4a90e2;
        box-shadow: 0 0 0 3px rgba(74, 144, 226, 0.2);
    }

    #month {
        min-width: 130px;
    }

    #year {
        min-width: 90px;
    }
</style>