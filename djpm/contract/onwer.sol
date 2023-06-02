
interface AspectOwnable {
    function isOwner(address sender) external view returns (bool result);
}