import utils.network as network


def main():
    print(network.network_available('https://www.youtube.com', 6000))


if __name__ == '__main__':
    main()
