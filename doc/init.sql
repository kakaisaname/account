# --初始化系统红包账户
INSERT INTO `account`(`id`, `account_no`, `account_name`, `account_type`, `currency_code`, `user_id`, `username`, `balance`, `status`, `created_at`, `updated_at`) VALUES (32937, '10000020190101010000000000000001', '系统红包账户', 2, 'CNY', '100001', '系统红包账户', 0.000000, 1, '2019-05-01 08:41:10.346', '2019-05-12 09:37:55.462');
INSERT INTO `account_log`(`id`, `trade_no`, `log_no`, `account_no`, `user_id`, `username`, `target_account_no`, `target_user_id`, `target_username`, `amount`, `balance`, `change_type`, `change_flag`, `status`, `decs`, `created_at`) VALUES (43208, '20190501084054283000000002110000', '20190501084054283000000002110000', '10000020190101010000000000000001', '100001', '系统红包账户', '10000020190101010000000000000001', '100001', '系统红包账户', 0.000000, 0.000000, 0, 0, 0, '开户', '2019-05-01 08:41:10.371');

