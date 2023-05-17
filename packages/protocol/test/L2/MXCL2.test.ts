import { expect } from "chai";
import { ethers } from "hardhat";
import { MXCL2 } from "../../typechain";
import deployAddressManager from "../utils/addressManager";
import { randomBytes32 } from "../utils/bytes";
import { deployMXCL2 } from "../utils/mxcL2";

describe("MXCL2", function () {
    let MXCL2: MXCL2;

    beforeEach(async function () {
        const signer = (await ethers.getSigners())[0];
        const addressManager = await deployAddressManager(signer);
        MXCL2 = await deployMXCL2(signer, addressManager);
    });

    describe("anchor()", async function () {
        it("should revert since ancestor hashes not written", async function () {
            await expect(
                MXCL2.anchor(Math.ceil(Math.random() * 1024), randomBytes32())
            ).to.be.revertedWith("L2_PUBLIC_INPUT_HASH_MISMATCH()");
        });
    });

    describe("getLatestSyncedHeader()", async function () {
        it("should be 0 because no headers have been synced", async function () {
            const hash = await MXCL2.getLatestSyncedHeader();
            expect(hash).to.be.eq(ethers.constants.HashZero);
        });
    });

    describe("getSyncedHeader()", async function () {
        it("should be 0 because header number has not been synced", async function () {
            const hash = await MXCL2.getSyncedHeader(1);
            expect(hash).to.be.eq(ethers.constants.HashZero);
        });
    });
});