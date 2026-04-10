const express = require('express');
const app = express();
const PORT = 3000;

const parse = (a, b) => [Number(a), Number(b)];

// Add
app.get('/add/:a/:b', (req, res) => {
  const [a, b] = parse(req.params.a, req.params.b);
  res.json({ calculation: `${a}+${b}`, result: a + b });
});

// Subtract
app.get('/substract/:a/:b', (req, res) => {
  const [a, b] = parse(req.params.a, req.params.b);
  res.json({ calculation: `${a}-${b}`, result: a - b });
});

// Multiply
app.get('/multiply/:a/:b', (req, res) => {
  const [a, b] = parse(req.params.a, req.params.b);
  res.json({ calculation: `${a}*${b}`, result: a * b });
});

// Divide
app.get('/divide/:a/:b', (req, res) => {
  const [a, b] = parse(req.params.a, req.params.b);

  if (b === 0) {
    return res.status(409).json({ error: "Division by zero" });
  }

  res.json({ calculation: `${a}/${b}`, result: a / b });
});

app.listen(PORT, () => {
  console.log(`Server running at http://localhost:${PORT}`);
});
