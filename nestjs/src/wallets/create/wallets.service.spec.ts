import { Test, TestingModule } from '@nestjs/testing';
import { WalletsCreateService } from './wallets.create.service';

describe('WalletsService', () => {
  let service: WalletsCreateService;

  beforeEach(async () => {
    const module: TestingModule = await Test.createTestingModule({
      providers: [WalletsCreateService],
    }).compile();

    service = module.get<WalletsCreateService>(WalletsCreateService);
  });

  it('should be defined', () => {
    expect(service).toBeDefined();
  });
});
