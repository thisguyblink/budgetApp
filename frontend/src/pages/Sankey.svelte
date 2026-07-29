<script>
    // BudgetSankey.svelte — Svelte 4 (legacy) syntax
    // npm install d3 d3-sankey

    import { onMount } from 'svelte';
    import { sankey, sankeyLinkHorizontal, sankeyJustify } from 'd3-sankey';

    /**
     * data shape:
     * {
     *   nodes: [{ id: 'Salary' }, { id: 'Housing' }, ...],
     *   links: [{ source: 'Salary', target: 'Housing', value: 1800 }, ...]
     * }
     *
     * In a Wails app you'll typically get this from your Go backend, e.g.:
     *   import { GetBudget } from '../../wailsjs/go/main/App';
     *   let raw = await GetBudget();
     * then pass it in as the `data` prop.
     */
    export let data = defaultData();
    export let currency = '$';

    let container;
    let width = 760;
    let height = 420;

    const nodeWidth = 16;
    const nodePadding = 24;

    // Recompute layout whenever data or width/height changes
    $: layout = computeLayout(data, width, height);

    function computeLayout(source, w, h) {
        if (!source || !source.nodes || !source.nodes.length) {
            return { nodes: [], links: [] };
        }

        const idIndex = new Map(source.nodes.map((n, i) => [n.id, i]));
        const graph = {
            nodes: source.nodes.map((n) => ({ ...n })),
            links: source.links.map((l) => ({
                source: idIndex.get(l.source),
                target: idIndex.get(l.target),
                value: l.value
            }))
        };

        const gen = sankey()
            .nodeId((_, i) => i)
            .nodeWidth(nodeWidth)
            .nodePadding(nodePadding)
            .nodeAlign(sankeyJustify)
            .extent([
                [1, 5],
                [w - 1, h - 5]
            ]);

        return gen(graph);
    }

    function color(name) {
        const palette = [
            '#2f6f4f', '#3d8b62', '#5aa876', '#7fb894',
            '#b08a2e', '#c9a13f', '#8a4a3d', '#a8624f',
            '#3d5a80', '#5d7ea3'
        ];
        let h = 0;
        for (const c of name) h = (h * 31 + c.charCodeAt(0)) >>> 0;
        return palette[h % palette.length];
    }

    function fmt(v) {
        return currency + Math.round(v).toLocaleString();
    }

    function linkPath(l) {
        return sankeyLinkHorizontal()(l);
    }

    onMount(() => {
        if (!container) return;
        const ro = new ResizeObserver((entries) => {
            for (const entry of entries) {
                width = Math.max(320, entry.contentRect.width);
            }
        });
        ro.observe(container);
        return () => ro.disconnect();
    });

    function defaultData() {
        return {
            nodes: [
                { id: 'Salary' }, { id: 'Freelance' },
                { id: 'Income' },
                { id: 'Housing' }, { id: 'Food' }, { id: 'Transport' },
                { id: 'Savings' }, { id: 'Uncategorized' },
                { id: 'Rent' }, { id: 'Utilities' },
                { id: 'Groceries' }, { id: 'Dining Out' }
            ],
            links: [
                { source: 'Salary', target: 'Income', value: 5200 },
                { source: 'Freelance', target: 'Income', value: 800 },
                { source: 'Income', target: 'Housing', value: 1800 },
                { source: 'Income', target: 'Food', value: 700 },
                { source: 'Income', target: 'Transport', value: 400 },
                { source: 'Income', target: 'Savings', value: 1500 },
                { source: 'Income', target: 'Uncategorized', value: 1600 },
                { source: 'Housing', target: 'Rent', value: 1500 },
                { source: 'Housing', target: 'Utilities', value: 300 },
                { source: 'Food', target: 'Groceries', value: 450 },
                { source: 'Food', target: 'Dining Out', value: 250 }
            ]
        };
    }
</script>

<div class="sankey-wrap" bind:this={container}>
    <svg {width} {height} viewBox="0 0 {width} {height}">
        <defs>
            {#each layout.links as l, i}
                <linearGradient id="grad-{i}" gradientUnits="userSpaceOnUse"
                                x1={l.source.x1} x2={l.target.x0}>
                    <stop offset="0%" stop-color={color(l.source.id)} stop-opacity="0.55" />
                    <stop offset="100%" stop-color={color(l.target.id)} stop-opacity="0.55" />
                </linearGradient>
            {/each}
        </defs>

        <g fill="none">
            {#each layout.links as l, i}
                <path
                        d={linkPath(l)}
                        stroke="url(#grad-{i})"
                        stroke-width={Math.max(1, l.width)}
                        class="link"
                >
                    <title>{l.source.id} → {l.target.id}: {fmt(l.value)}</title>
                </path>
            {/each}
        </g>

        <g>
            {#each layout.nodes as n}
                <rect
                        x={n.x0}
                        y={n.y0}
                        width={n.x1 - n.x0}
                        height={n.y1 - n.y0}
                        fill={color(n.id)}
                        rx="2"
                        class="node"
                >
                    <title>{n.id} {fmt(n.value)}</title>
                </rect>
                <text
                        x={n.x0 < width / 2 ? n.x1 + 8 : n.x0 - 8}
                        y={(n.y0 + n.y1) / 2}
                        dy="0.32em"
                        text-anchor={n.x0 < width / 2 ? 'start' : 'end'}
                        class="label"
                >
                    {n.id}:
                    <tspan class="value">{fmt(n.value)}</tspan>
                </text>
            {/each}
        </g>
    </svg>
</div>

<style>
    .sankey-wrap {
        width: 100%;
        font-family: system-ui, -apple-system, sans-serif;
    }

    .link {
        transition: stroke-opacity 0.15s ease;
        mix-blend-mode: multiply;
    }
    .link:hover {
        stroke-opacity: 0.85;
    }

    .node {
        stroke: rgba(0, 0, 0, 0.15);
        stroke-width: 1;
    }

    .label {
        font-size: 18px;
        fill: #7fb894;
    }
    .value {
        font-size: 16px;
        fill: #3d8b62;
    }
    .value::before {
        content: '  ·  ';
    }
</style>