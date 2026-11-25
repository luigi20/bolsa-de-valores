import { Module } from '@nestjs/common';
import { WalletAssetsGetAllController } from './get_all/wallet_assets.get_all.controller';
import { WalletAssetsCreateController } from './create/wallet_assets.create.controller';
import { WalletAssetsGetAllService } from './get_all/wallet_assets.get_all.service';
import { WalletAssetsCreateService } from './create/wallet_assets.create.service';

@Module({
  controllers: [WalletAssetsCreateController, WalletAssetsGetAllController],
  providers: [WalletAssetsCreateService, WalletAssetsGetAllService],
})
export class WalletAssetsModule {}
