import pandas as pd
import plotly.express
import plotly.offline
import plotly.io

# Read the benchmark results
df = pd.read_csv("bench.csv")

# Plot the results
fig = plotly.express.bar(
    df,
    x="name",
    y="ns_per_op",
    color="name",
    title="Benchmark Results",
    labels={"name": "Benchmark", "ns_per_op": "Time (s)"},
    template="plotly_dark",
)

# Save the plot
plotly.offline.plot(fig, filename="benchmark.html", auto_open=False)
plotly.io.write_image(fig, "benchmark.png", width=1280, height=720, scale=2)
