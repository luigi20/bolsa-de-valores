import { Body, Controller, Param, Post } from '@nestjs/common';
import { WalletAssetsCreateService } from './wallet_assets.create.service';

@Controller('wallet_assets/:wallet_id/assets')
export class WalletAssetsCreateController {
  constructor(
    private readonly walletAssetsCreateService: WalletAssetsCreateService,
  ) {}

  @Post()
  public async create(
    @Param('wallet_id') wallet_id: string,
    @Body() body: { asset_id: string; shares: number },
  ) {
    return await this.walletAssetsCreateService.execute({
      wallet_id: wallet_id,
      asset_id: body.asset_id,
      shares: body.shares,
    });
  }
}
