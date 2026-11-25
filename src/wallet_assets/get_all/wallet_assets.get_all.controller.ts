import { Controller, Get, Param } from '@nestjs/common';
import { WalletAssetsGetAllService } from './wallet_assets.get_all.service';

@Controller('wallet_assets/:wallet_id/assets')
export class WalletAssetsGetAllController {
  constructor(
    private readonly walletAssetsGetAllService: WalletAssetsGetAllService,
  ) {}

  @Get()
  public async getAll(@Param('wallet_id') wallet_id: string) {
    return await this.walletAssetsGetAllService.execute(wallet_id);
  }
}
