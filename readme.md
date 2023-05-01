# Review Rollup (RevUp)
**Review Rollup (RevUp)** is a blockchain built with [Ignite CLI](https://ignite.com/cli).

## Get started

**Review Rollup (RevUp)** is a decentralized application (dApp) running on the Blockspacerace, a testnet network for Celestia, which allows users to submit their opinions on specific products, rate them, and view the comments made by others. You can use the ready-to-use setup script or follow the manual installation steps to get started.

### Details

After the installation process, you can use the following commands to use the application.

```
revupd tx revup create-review product rate user_comment --from revup-key --keyring-backend test
```

Let's examine in-depth:

You can specify the product name with the ```product``` parameter.
```rate``` indicates the score you will give to the product. You can rate it between 1 and 5.
The ```user_comment``` parameter contains your detailed comment about the product. You can write your comment inside double quotes ("").


To view the recorded comments on the chain, you can use the following command.

```
revupd query revup comments
```

### Web UI

With a Web UI , you can allow users to add comments to the chain and view the comments. For this, I have prepared a simple panel using Flask, and it looks like this.

Using the example UI, you can do the following:

- You can predefine the products users will comment on.
- You can limit the duration users can make comments and prevent congestion on the network.
- You can determine the comment scoring and allow the user to give a score between 1 and 5.
- You can view previously made comments with a single click.

What you can do is limited only by your imagination!


### Install
For the installation steps, you can visit the install page to obtain detailed information.


## Learn more

- [Ignite CLI](https://ignite.com/cli)
- [Tutorials](https://docs.ignite.com/guide)
- [Ignite CLI docs](https://docs.ignite.com)
- [Cosmos SDK docs](https://docs.cosmos.network)
- [Developer Chat](https://discord.gg/ignite)
