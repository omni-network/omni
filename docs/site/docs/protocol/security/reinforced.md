---
sidebar_position: 2
---

# Reinforced Security

Omni achieves stronger and more stable security guarantees than existing interoperability protocols by deriving its cryptoeconomic security from restaked **\$ETH**. Omni extends its security model further by incorporating staked **\$OMNI** using a dual staking model. Effectively, the total cryptoeconomic security of Omni is determined by the combined value of restaked **\$ETH** and staked **\$OMNI**.

Using this dual staking model, the total cryptoeconomic security $\textit{C}$ of the system is given by the formula:

<br/>
<div style={{ textAlign: 'center', fontSize: '1.7em' }}>
$C = \frac{2}{3} \sum_{a=0}^{m} \sum_{v=0}^{n} P_a(S_{a,v})$
</div>
<br/>

where:

- $S_{a,v}$ is the amount staked by validator $v$ for asset $a$
- $P_a$ is the function mapping the amount of asset $a$ staked to validator power
- $n$ is the total number of validators
- $m$ is the total number of unique staked asset types

Omni uses this reinforced staking model to scale its security across two dimensions. Restaked **\$ETH** anchors Omni’s security to Ethereum L1, enabling it to grow in line with Ethereum’s own security budget. The addition of staked **\$OMNI** builds upon this base, expanding Omni’s security alongside its own network activity. Collectively, these two complementary mechanisms provide robust and dynamic security guarantees for Omni, setting a new standard for secure interoperability for the Ethereum ecosystem.
