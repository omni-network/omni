import { parseEther } from "viem";

// Helper function to convert ether amount to hex string
const etherToHex = (amount: string) => "0x" + parseEther(amount).toString(16);

// Helper function to make API calls
async function makeApiCall(endpoint: string, params: any) {
  const response = await fetch(
    `https://solver.staging.omni.network/api/v1/${endpoint}`,
    {
      method: "POST",
      headers: {
        "Content-Type": "application/json",
      },
      body: JSON.stringify(params),
    }
  );

  if (!response.ok) {
    const errorText = await response.text();
    throw new Error(
      `${endpoint} error! status: ${response.status}, body: ${errorText}`
    );
  }

  return response.json();
}

// Helper function to build and check an order from a quote
async function buildAndCheckOrder(quote: any) {
  const callerAddress = "0xA779fC675Db318dab004Ab8D538CB320D0013F42";

  // Build order params
  const orderParams = {
    sourceChainId: 17000,
    destChainId: 84532,
    owner: callerAddress,
    deposit: {
      isNative: true,
      amount: quote.deposit.amount,
    },
    calls: [
      {
        target: callerAddress,
        value: quote.expense.amount,
      },
    ],
    expense: {
      isNative: true,
      amount: quote.expense.amount,
    },
  };

  // Check order and return result
  const checkResult = await makeApiCall("check", orderParams);
  console.log("Order Check Response:", JSON.stringify(checkResult, null, 2));
  return checkResult;
}

async function testExpenseQuote() {
  try {
    // Build quote request
    const params = {
      sourceChainId: 17000, // Ethereum Holesky
      destChainId: 84532, // Base Sepolia
      mode: "expense",
      deposit: {
        isNative: true,
        amount: etherToHex("0.01"),
      },
      expense: {
        isNative: true,
      },
      enabled: true,
    };

    // Get quote and build order
    const quote = await makeApiCall("quote", params);
    console.log("Expense Quote Response:", JSON.stringify(quote, null, 2));
    await buildAndCheckOrder(quote);
  } catch (error) {
    console.error("Error:", error);
  }
}

async function testDepositQuote() {
  try {
    // Build quote request
    const params = {
      sourceChainId: 17000, // Ethereum Holesky
      destChainId: 84532, // Base Sepolia
      mode: "deposit",
      deposit: {
        isNative: true,
      },
      expense: {
        isNative: true,
        amount: etherToHex("0.01"),
      },
      enabled: true,
    };

    // Get quote and build order
    const quote = await makeApiCall("quote", params);
    console.log("Deposit Quote Response:", JSON.stringify(quote, null, 2));
    await buildAndCheckOrder(quote);
  } catch (error) {
    console.error("Error:", error);
  }
}

await testExpenseQuote();
console.log("");
await testDepositQuote();
