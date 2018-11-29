def manual_iter():
	with open('/etc/passwd') as f:
		try:
			while True:
				line = next(f)
				print(line,end='')
		except StopIteration:
			pass

manual_iter()

